// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

/************************************************
  generated by IDE. for [InternetAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *InternetAPI) Reset() *InternetAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *InternetAPI) Offset(offset int) *InternetAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *InternetAPI) Limit(limit int) *InternetAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *InternetAPI) Include(key string) *InternetAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *InternetAPI) Exclude(key string) *InternetAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *InternetAPI) FilterBy(key string, value interface{}) *InternetAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *InternetAPI) FilterMultiBy(key string, value interface{}) *InternetAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *InternetAPI) WithNameLike(name string) *InternetAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *InternetAPI) WithTag(tag string) *InternetAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *InternetAPI) WithTags(tags []string) *InternetAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *InternetAPI) WithSizeGib(size int) *InternetAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *InternetAPI) WithSharedScope() *InternetAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *InternetAPI) WithUserScope() *InternetAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *InternetAPI) SortBy(key string, reverse bool) *InternetAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *InternetAPI) SortByName(reverse bool) *InternetAPI {
	api.sortByName(reverse)
	return api
}

// func (api *InternetAPI) SortBySize(reverse bool) *InternetAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *InternetAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *InternetAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *InternetAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *InternetAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *InternetAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルター
func (api *InternetAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *InternetAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

// SetNameLike 名称条件
func (api *InternetAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

// SetTag タグ条件
func (api *InternetAPI) SetTag(tag string) {
	api.FilterBy("Tags.Name", tag)
}

// SetTags タグ(複数)条件
func (api *InternetAPI) SetTags(tags []string) {
	api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *InternetAPI) SetSizeGib(size int)  {
// 	api.FilterBy("SizeMB", size*1024)
// }

// func (api *InternetAPI) SetSharedScope()  {
// 	api.FilterBy("Scope", "shared")
// }

// func (api *InternetAPI) SetUserScope()  {
// 	api.FilterBy("Scope", "user")
// }

// SetSortBy 指定キーでのソート
func (api *InternetAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

// SetSortByName 名称でのソート
func (api *InternetAPI) SetSortByName(reverse bool) {
	api.sortByName(reverse)
}

// func (api *InternetAPI) SetSortBySize(reverse bool)  {
// 	api.sortBy("SizeMB", reverse)
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// New 新規作成用パラメーター作成
func (api *InternetAPI) New() *sacloud.Internet {
	return &sacloud.Internet{}
}

// Create 新規作成
func (api *InternetAPI) Create(value *sacloud.Internet) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

// Read 読み取り
func (api *InternetAPI) Read(id int64) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *InternetAPI) Update(id int64, value *sacloud.Internet) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// Delete 削除
func (api *InternetAPI) Delete(id int64) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *InternetAPI) setStateValue(setFunc func(*sacloud.Request)) *InternetAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *InternetAPI) request(f func(*sacloud.Response) error) (*sacloud.Internet, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Internet, nil
}

func (api *InternetAPI) createRequest(value *sacloud.Internet) *sacloud.Request {
	req := &sacloud.Request{}
	req.Internet = value
	return req
}
