package builder

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"time"
)

/**********************************************************
  Type : DiskBuildEvents
**********************************************************/

// DiskBuildEvents ディスク構築時イベント種別
type DiskBuildEvents int

const (
	// DiskBuildOnStart ディスク作成 開始
	DiskBuildOnStart DiskBuildEvents = iota

	// DiskBuildOnCreateSSHKeyBefore SSHキー作成 開始時
	DiskBuildOnCreateSSHKeyBefore

	// DiskBuildOnCreateSSHKeyAfter SSHキー作成 終了時
	DiskBuildOnCreateSSHKeyAfter

	// DiskBuildOnCreateNoteBefore スタートアップスクリプト作成 開始時
	DiskBuildOnCreateNoteBefore

	// DiskBuildOnCreateNoteAfter スタートアップスクリプト作成 終了時
	DiskBuildOnCreateNoteAfter

	// DiskBuildOnCreateDiskBefore ディスク作成 開始時
	DiskBuildOnCreateDiskBefore

	// DiskBuildOnCreateDiskAfter ディスク作成 終了時
	DiskBuildOnCreateDiskAfter

	// DiskBuildOnEditDiskBefore ディスク修正 開始時
	DiskBuildOnEditDiskBefore

	// DiskBuildOnEditDiskAfter ディスク修正 終了時
	DiskBuildOnEditDiskAfter

	// DiskBuildOnCleanupSSHKeyBefore SSHキークリーンアップ 開始時
	DiskBuildOnCleanupSSHKeyBefore

	// DiskBuildOnCleanupSSHKeyAfter SSHキークリーンアップ 終了時
	DiskBuildOnCleanupSSHKeyAfter

	// DiskBuildOnCleanupNoteBefore スタートアップスクリプトクリーンアップ 開始時
	DiskBuildOnCleanupNoteBefore

	// DiskBuildOnCleanupNoteAfter スタートアップスクリプトクリーンアップ 終了時
	DiskBuildOnCleanupNoteAfter

	// DiskBuildOnComplete ディスク作成 完了
	DiskBuildOnComplete
)

// DiskBuildEventHandler ディスク構築時イベントハンドラ型
type DiskBuildEventHandler func(value *DiskBuildValue, result *DiskBuildResult)

/**********************************************************
  Type : DiskBuilder
**********************************************************/

// DiskBuilder ディスクビルダー
type DiskBuilder struct {
	*baseBuilder
	buildEventHandlers map[DiskBuildEvents]DiskBuildEventHandler

	name            string
	size            int
	distantFrom     []int64
	planID          sacloud.DiskPlanID
	connection      sacloud.EDiskConnection
	sourceArchiveID int64
	sourceDiskID    int64
	description     string
	*sacloud.TagsType
	iconID int64

	ipAddress          string
	networkMaskLen     int
	defaultRoute       string
	password           string
	hostName           string
	disablePWAuth      bool
	sshKeys            []string
	isSSHKeysEphemeral bool
	sshKeyIDs          []int64
	notes              []string
	isNotesEphemeral   bool
	noteIDs            []int64

	currentDiskBuildValue  *DiskBuildValue
	currentDiskBuildResult *DiskBuildResult
}

const (
	// DefaultDiskPlanID ディスクプラン(デフォルト値)
	DefaultDiskPlanID = sacloud.DiskPlanSSDID
	// DefaultDiskConnection ディスク接続方法(デフォルト値)
	DefaultDiskConnection = sacloud.DiskConnectionVirtio
	// DefaultDiskSize ディスクサイズ(デフォルト値)
	DefaultDiskSize = 20
	// DefaultDiskIsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ(デフォルト値)
	DefaultDiskIsSSHKeysEphemeral = true
	// DefaultDiskIsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ(デフォルト値)
	DefaultDiskIsNotesEphemeral = true
)

// Disk ディスクビルダーの作成
func Disk(client *api.Client, name string) *DiskBuilder {
	return &DiskBuilder{
		baseBuilder: &baseBuilder{
			client: client,
		},
		buildEventHandlers: map[DiskBuildEvents]DiskBuildEventHandler{},
		name:               name,
		size:               DefaultDiskSize,
		planID:             DefaultDiskPlanID,
		connection:         DefaultDiskConnection,
		TagsType:           &sacloud.TagsType{},
		isSSHKeysEphemeral: DefaultDiskIsSSHKeysEphemeral,
		isNotesEphemeral:   DefaultDiskIsNotesEphemeral,
	}
}

// GetName ディスク名 取得
func (b *DiskBuilder) GetName() string {
	return b.name
}

// SetName ディスク名 設定
func (b *DiskBuilder) SetName(name string) *DiskBuilder {
	b.name = name
	return b
}

// GetSize ディスクサイズ(GB単位) 取得
func (b *DiskBuilder) GetSize() int {
	return b.size
}

// SetSize ディスクサイズ(GB単位) 設定
func (b *DiskBuilder) SetSize(size int) *DiskBuilder {
	b.size = size
	return b
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *DiskBuilder) GetDistantFrom() []int64 {
	return b.distantFrom
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *DiskBuilder) SetDistantFrom(diskIDs []int64) *DiskBuilder {
	b.distantFrom = diskIDs
	return b
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *DiskBuilder) AddDistantFrom(diskID int64) *DiskBuilder {
	b.distantFrom = append(b.distantFrom, diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *DiskBuilder) ClearDistantFrom() *DiskBuilder {
	b.distantFrom = []int64{}
	return b
}

// GetPlanID ディスクプラン(SSD/HDD) 取得
func (b *DiskBuilder) GetPlanID() sacloud.DiskPlanID {
	return b.planID
}

// SetPlanID ディスクプラン(SSD/HDD) 設定
func (b *DiskBuilder) SetPlanID(planID sacloud.DiskPlanID) *DiskBuilder {
	b.planID = planID
	return b
}

// GetConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *DiskBuilder) GetConnection() sacloud.EDiskConnection {
	return b.connection
}

// SetConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *DiskBuilder) SetConnection(connection sacloud.EDiskConnection) *DiskBuilder {
	b.connection = connection
	return b
}

// GetSourceArchiveID ソースアーカイブID 取得
func (b *DiskBuilder) GetSourceArchiveID() int64 {
	return b.sourceArchiveID
}

// SetSourceArchiveID ソースアーカイブID 設定
func (b *DiskBuilder) SetSourceArchiveID(id int64) *DiskBuilder {
	b.sourceArchiveID = id
	b.sourceDiskID = 0
	return b
}

// GetSourceDiskID ソースディスクID 取得
func (b *DiskBuilder) GetSourceDiskID() int64 {
	return b.sourceDiskID
}

// SetSourceDiskID ソースディスクID 設定
func (b *DiskBuilder) SetSourceDiskID(id int64) *DiskBuilder {
	b.sourceArchiveID = 0
	b.sourceDiskID = id
	return b
}

// GetDescription 説明 取得
func (b *DiskBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *DiskBuilder) SetDescription(desc string) *DiskBuilder {
	b.description = desc
	return b
}

// GetTags タグ 取得
func (b *DiskBuilder) GetTags() []string {
	return b.Tags
}

// SetTags タグ 設定
func (b *DiskBuilder) SetTags(tags []string) *DiskBuilder {
	b.Tags = tags
	return b
}

// GetIconID アイコンID 取得
func (b *DiskBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *DiskBuilder) SetIconID(id int64) *DiskBuilder {
	b.iconID = id
	return b
}

// GetIPAddress IPアドレス 取得
func (b *DiskBuilder) GetIPAddress() string {
	return b.ipAddress
}

// SetIPAddress IPアドレス 設定
func (b *DiskBuilder) SetIPAddress(ip string) *DiskBuilder {
	b.ipAddress = ip
	return b
}

// GetNetworkMaskLen ネットワークマスク長 取得
func (b *DiskBuilder) GetNetworkMaskLen() int {
	return b.networkMaskLen
}

// SetNetworkMaskLen ネットワークマスク長 設定
func (b *DiskBuilder) SetNetworkMaskLen(masklen int) *DiskBuilder {
	b.networkMaskLen = masklen
	return b
}

// GetDefaultRoute デフォルトルート 取得
func (b *DiskBuilder) GetDefaultRoute() string {
	return b.defaultRoute
}

// SetDefaultRoute デフォルトルート 設定
func (b *DiskBuilder) SetDefaultRoute(route string) *DiskBuilder {
	b.defaultRoute = route
	return b
}

// GetPassword パスワード 取得
func (b *DiskBuilder) GetPassword() string {
	return b.password
}

// SetPassword パスワード 設定
func (b *DiskBuilder) SetPassword(password string) *DiskBuilder {
	b.password = password
	return b
}

// GetHostName ホスト名 取得
func (b *DiskBuilder) GetHostName() string {
	return b.hostName
}

// SetHostName ホスト名 設定
func (b *DiskBuilder) SetHostName(name string) *DiskBuilder {
	b.hostName = name
	return b
}

// IsDisablePWAuth パスワード認証無効化フラグ 取得
func (b *DiskBuilder) IsDisablePWAuth() bool {
	return b.disablePWAuth
}

// SetDisablePWAuth パスワード認証無効化フラグ 設定
func (b *DiskBuilder) SetDisablePWAuth(disable bool) *DiskBuilder {
	b.disablePWAuth = disable
	return b
}

// AddSSHKeyID 公開鍵ID 追加
func (b *DiskBuilder) AddSSHKeyID(sshKeyID int64) *DiskBuilder {
	b.sshKeyIDs = append(b.sshKeyIDs, sshKeyID)
	return b
}

// ClearSSHKeyIDs 公開鍵ID クリア
func (b *DiskBuilder) ClearSSHKeyIDs() *DiskBuilder {
	b.sshKeyIDs = []int64{}
	return b
}

// GetSSHKeyIds 公開鍵ID 取得
func (b *DiskBuilder) GetSSHKeyIds() []int64 {
	return b.sshKeyIDs
}

// AddSSHKey 公開鍵 追加
func (b *DiskBuilder) AddSSHKey(sshKey string) *DiskBuilder {
	b.sshKeys = append(b.sshKeys, sshKey)
	return b
}

// ClearSSHKey 公開鍵 クリア
func (b *DiskBuilder) ClearSSHKey() *DiskBuilder {
	b.sshKeys = []string{}
	return b
}

// GetSSHKeys 公開鍵 取得
func (b *DiskBuilder) GetSSHKeys() []string {
	return b.sshKeys
}

// AddNote スタートアップスクリプト 追加
func (b *DiskBuilder) AddNote(note string) *DiskBuilder {
	b.notes = append(b.notes, note)
	return b
}

// ClearNotes スタートアップスクリプト クリア
func (b *DiskBuilder) ClearNotes() *DiskBuilder {
	b.notes = []string{}
	return b
}

// GetNotes スタートアップスクリプト 取得
func (b *DiskBuilder) GetNotes() []string {
	return b.notes
}

// AddNoteID スタートアップスクリプトID 追加
func (b *DiskBuilder) AddNoteID(noteID int64) *DiskBuilder {
	b.noteIDs = append(b.noteIDs, noteID)
	return b
}

// ClearNoteIDs スタートアップスクリプトID クリア
func (b *DiskBuilder) ClearNoteIDs() *DiskBuilder {
	b.noteIDs = []int64{}
	return b
}

// GetNoteIDs スタートアップスクリプトID 取得
func (b *DiskBuilder) GetNoteIDs() []int64 {
	return b.noteIDs
}

// IsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 取得
func (b *DiskBuilder) IsSSHKeysEphemeral() bool {
	return b.isSSHKeysEphemeral
}

// SetSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *DiskBuilder) SetSSHKeysEphemeral(isEphemeral bool) *DiskBuilder {
	b.isSSHKeysEphemeral = isEphemeral
	return b
}

// IsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 取得
func (b *DiskBuilder) IsNotesEphemeral() bool {
	return b.isNotesEphemeral
}

// SetNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *DiskBuilder) SetNotesEphemeral(isEphemeral bool) *DiskBuilder {
	b.isNotesEphemeral = isEphemeral
	return b
}

// SetEventHandler イベントハンドラ 登録
func (b *DiskBuilder) SetEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) *DiskBuilder {
	b.buildEventHandlers[event] = handler
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *DiskBuilder) ClearEventHandler(event DiskBuildEvents) *DiskBuilder {
	delete(b.buildEventHandlers, event)
	return b
}

// GetEventHandler イベントハンドラ取得
func (b *DiskBuilder) GetEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// Build ディスクの構築
func (b *DiskBuilder) Build() (*DiskBuildResult, error) {

	// start
	b.callEventHandlerIfExists(DiskBuildOnStart)

	// create parameters
	b.currentDiskBuildValue = &DiskBuildValue{}
	b.currentDiskBuildResult = &DiskBuildResult{}

	if err := b.buildDiskParams(); err != nil {
		return nil, err
	}

	// create disk
	b.callEventHandlerIfExists(DiskBuildOnCreateDiskBefore)
	if err := b.createDisk(b.currentDiskBuildValue.Disk); err != nil {
		return b.currentDiskBuildResult, err
	}
	b.callEventHandlerIfExists(DiskBuildOnCreateDiskAfter)

	// edit disk
	if b.currentDiskBuildValue.Edit != nil {
		b.callEventHandlerIfExists(DiskBuildOnEditDiskBefore)
		if err := b.editDisk(b.currentDiskBuildValue.Edit); err != nil {
			return b.currentDiskBuildResult, err
		}
		b.callEventHandlerIfExists(DiskBuildOnEditDiskAfter)
	}
	// cleanup
	if b.isSSHKeysEphemeral && len(b.currentDiskBuildResult.SSHKeys) > 0 {
		b.callEventHandlerIfExists(DiskBuildOnCleanupSSHKeyBefore)

		for _, key := range b.currentDiskBuildResult.SSHKeys {
			_, err := b.client.SSHKey.Delete(key.ID)
			if err != nil {
				return b.currentDiskBuildResult, err
			}
		}

		b.callEventHandlerIfExists(DiskBuildOnCleanupSSHKeyAfter)
	}
	if b.isNotesEphemeral && len(b.currentDiskBuildResult.Notes) > 0 {
		b.callEventHandlerIfExists(DiskBuildOnCleanupNoteBefore)

		for _, note := range b.currentDiskBuildResult.Notes {
			_, err := b.client.Note.Delete(note.ID)
			if err != nil {
				return b.currentDiskBuildResult, err
			}
		}

		b.callEventHandlerIfExists(DiskBuildOnCleanupNoteAfter)
	}

	b.callEventHandlerIfExists(DiskBuildOnComplete)

	return b.currentDiskBuildResult, nil
}

func (b *DiskBuilder) buildDiskParams() error {
	if err := b.buildDiskParam(); err != nil {
		return err
	}

	if !b.isNeedDiskEdit() {
		return nil
	}

	if err := b.buildDiskEditParam(); err != nil {
		return err
	}
	return nil
}

func (b *DiskBuilder) buildDiskParam() error {
	v := b.currentDiskBuildValue

	v.Disk = b.client.Disk.New()
	d := v.Disk
	d.Name = b.name
	d.SizeMB = b.size * 1024
	d.DistantFrom = b.distantFrom
	d.Plan = b.planID.ToResource()
	d.Connection = b.connection
	d.Description = b.description
	d.Tags = b.GetTags()
	if b.iconID > 0 {
		d.Icon = &sacloud.Icon{Resource: sacloud.NewResource(b.iconID)}
	}

	if b.sourceArchiveID > 0 {
		d.SetSourceArchive(b.sourceArchiveID)
		d.SourceDisk = nil
	}
	if b.sourceDiskID > 0 {
		d.SourceArchive = nil
		d.SetSourceDisk(b.sourceDiskID)
	}

	return nil
}

func (b *DiskBuilder) buildDiskEditParam() error {
	v := b.currentDiskBuildValue

	// for DiskEditValue( POST /disk/config )
	v.Edit = b.client.Disk.NewCondig()
	e := v.Edit
	if b.ipAddress != "" {
		e.SetUserIPAddress(b.ipAddress)
	}
	if b.networkMaskLen > 0 {
		e.SetNetworkMaskLen(fmt.Sprintf("%d", b.networkMaskLen))
	}
	if b.defaultRoute != "" {
		e.SetDefaultRoute(b.defaultRoute)
	}
	if b.password != "" {
		e.SetPassword(b.password)
	}
	if b.hostName != "" {
		e.SetHostName(b.hostName)
	}
	e.SetDisablePWAuth(b.disablePWAuth)

	sshKeyIDs := []string{}
	if len(b.sshKeyIDs) > 0 {
		sshKeyIDs = append(sshKeyIDs, b.getStrSSHKeyIDs()...)
	}
	if len(b.sshKeys) > 0 {

		createdIDs, err := b.createSSHKeys()
		if err != nil {
			return err
		}
		sshKeyIDs = append(sshKeyIDs, createdIDs...)
	}
	if len(sshKeyIDs) > 0 {
		e.SetSSHKeys(sshKeyIDs)
	}

	noteIDs := []string{}
	if len(b.noteIDs) > 0 {
		noteIDs = append(noteIDs, b.getStrNoteIDs()...)
	}
	if len(b.notes) > 0 {
		createdIDs, err := b.createNotes()
		if err != nil {
			return err
		}
		noteIDs = append(noteIDs, createdIDs...)

	}
	e.SetNotes(noteIDs)

	return nil
}

func (b *DiskBuilder) createSSHKeys() ([]string, error) {
	createdIDs := []string{}
	for _, strKey := range b.sshKeys {
		key, err := b.createSSHKey(strKey)
		if err != nil {
			return createdIDs, err
		}
		createdIDs = append(createdIDs, key.GetStrID())
	}
	return createdIDs, nil
}

func (b *DiskBuilder) createSSHKey(strKey string) (*sacloud.SSHKey, error) {

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateSSHKeyBefore)

	keyReq := b.client.SSHKey.New()
	keyReq.Name = fmt.Sprintf("publickey-%s", time.Now())
	keyReq.PublicKey = strKey

	key, err := b.client.SSHKey.Create(keyReq)
	if err != nil {
		return nil, err
	}
	b.currentDiskBuildResult.addSSHKey(key)

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateSSHKeyAfter)

	return key, nil

}

func (b *DiskBuilder) createNotes() ([]string, error) {
	createdIDs := []string{}
	for _, strNote := range b.notes {
		note, err := b.createNote(strNote)
		if err != nil {
			return createdIDs, err
		}
		createdIDs = append(createdIDs, note.GetStrID())
	}
	return createdIDs, nil

}

func (b *DiskBuilder) createNote(strNote string) (*sacloud.Note, error) {

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateNoteBefore)

	noteReq := b.client.Note.New()
	noteReq.Name = fmt.Sprintf("note-%s", time.Now())
	noteReq.Content = strNote

	note, err := b.client.Note.Create(noteReq)
	if err != nil {
		return nil, err
	}
	b.currentDiskBuildResult.addNote(note)

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateNoteAfter)

	return note, nil

}

func (b *DiskBuilder) createDisk(diskReq *sacloud.Disk) error {
	disk, err := b.client.Disk.Create(diskReq)
	if err != nil {
		return err
	}

	b.currentDiskBuildResult.Disk = disk
	//wait
	if err := b.client.Disk.SleepWhileCopying(disk.ID, b.client.DefaultTimeoutDuration); err != nil {
		return err
	}

	return nil
}

func (b *DiskBuilder) editDisk(editReq *sacloud.DiskEditValue) error {
	_, err := b.client.Disk.Config(b.currentDiskBuildResult.Disk.ID, editReq)
	if err != nil {
		return err
	}
	return nil
}

func (b *DiskBuilder) isNeedDiskEdit() bool {
	if b.sourceArchiveID == 0 && b.sourceDiskID == 0 {
		// blank disk
		return false
	}

	return b.ipAddress != "" ||
		b.networkMaskLen > 0 ||
		b.defaultRoute != "" ||
		b.password != "" ||
		b.hostName != "" ||
		len(b.sshKeyIDs) > 0 ||
		len(b.sshKeys) > 0 ||
		len(b.noteIDs) > 0 ||
		len(b.notes) > 0
}

func (b *DiskBuilder) getStrSSHKeyIDs() []string {
	return b.toStringList(b.sshKeyIDs)
}

func (b *DiskBuilder) getStrNoteIDs() []string {
	return b.toStringList(b.noteIDs)
}

func (b *DiskBuilder) callEventHandlerIfExists(event DiskBuildEvents) {
	if handler, ok := b.buildEventHandlers[event]; ok {
		handler(b.currentDiskBuildValue, b.currentDiskBuildResult)
	}
}

/**********************************************************
  Type : DiskBuildValue
**********************************************************/

// DiskBuildValue ディスク構築用パラメータ
type DiskBuildValue struct {
	// Disk ディスク作成用パラメータ
	Disk *sacloud.Disk
	// Edit ディスクの編集用パラメータ
	Edit *sacloud.DiskEditValue
}

/**********************************************************
  Type : DiskBuildResult
**********************************************************/

// DiskBuildResult ディスク構築結果
type DiskBuildResult struct {
	// Disk ディスク
	Disk *sacloud.Disk
	// Notes スタートアップスクリプト
	Notes []*sacloud.Note
	// SSHKeys 公開鍵
	SSHKeys []*sacloud.SSHKey
}

func (d *DiskBuildResult) addNote(note *sacloud.Note) {
	d.Notes = append(d.Notes, note)
}

func (d *DiskBuildResult) addSSHKey(key *sacloud.SSHKey) {
	d.SSHKeys = append(d.SSHKeys, key)
}
