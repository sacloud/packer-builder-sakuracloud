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
  generated by IDE. for [IPAddressAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *IPAddressAPI) Reset() *IPAddressAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *IPAddressAPI) Offset(offset int) *IPAddressAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *IPAddressAPI) Limit(limit int) *IPAddressAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *IPAddressAPI) Include(key string) *IPAddressAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *IPAddressAPI) Exclude(key string) *IPAddressAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルタ
func (api *IPAddressAPI) FilterBy(key string, value interface{}) *IPAddressAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *IPAddressAPI) FilterMultiBy(key string, value interface{}) *IPAddressAPI {
	api.filterBy(key, value, true)
	return api
}

//func (api *IPAddressAPI) WithNameLike(name string) *IPAddressAPI {
//	return api.FilterBy("Name", name)
//}

//func (api *IPAddressAPI) WithTag(tag string) *IPAddressAPI {
//	return api.FilterBy("Tags.Name", tag)
//}
//func (api *IPAddressAPI) WithTags(tags []string) *IPAddressAPI {
//	return api.FilterBy("Tags.Name", []interface{}{tags})
//}
//
//func (api *IPAddressAPI) WithSizeGib(size int) *IPAddressAPI {
//	api.FilterBy("SizeMB", size*1024)
//	return api
//}
//
//func (api *IPAddressAPI) WithSharedScope() *IPAddressAPI {
//	api.FilterBy("Scope", "shared")
//	return api
//}
//
//func (api *IPAddressAPI) WithUserScope() *IPAddressAPI {
//	api.FilterBy("Scope", "user")
//	return api
//}

// SortBy 指定キーでのソート
func (api *IPAddressAPI) SortBy(key string, reverse bool) *IPAddressAPI {
	api.sortBy(key, reverse)
	return api
}

//// SortByName 名称でのソート
//func (api *IPAddressAPI) SortByName(reverse bool) *IPAddressAPI {
//	api.sortByName(reverse)
//	return api
//}

//func (api *IPAddressAPI) SortBySize(reverse bool) *IPAddressAPI {
//	api.sortBy("SizeMB", reverse)
//	return api
//}

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *IPAddressAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *IPAddressAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *IPAddressAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *IPAddressAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *IPAddressAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルタ
func (api *IPAddressAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *IPAddressAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

//func (api *IPAddressAPI) SetNameLike(name string)  {
//	api.FilterBy("Name", name)
//}

//func (api *IPAddressAPI) SetTag(tag string)  {
//	api.FilterBy("Tags.Name", tag)
//}
//func (api *IPAddressAPI) SetTags(tags []string)  {
//	api.FilterBy("Tags.Name", []interface{}{tags})
//}
//
//func (api *IPAddressAPI) SetSizeGib(size int)  {
//	api.FilterBy("SizeMB", size*1024)
//}
//
//func (api *IPAddressAPI) SetSharedScope()  {
//	api.FilterBy("Scope", "shared")
//}
//
//func (api *IPAddressAPI) SetUserScope()  {
//	api.FilterBy("Scope", "user")
//}

// SetSortBy 指定キーでのソート
func (api *IPAddressAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

//// SetSortByName 名称でのソート
//func (api *IPAddressAPI) SetSortByName(reverse bool)  {
//	api.sortByName(reverse)
//}

//func (api *IPAddressAPI) SetSortBySize(reverse bool)  {
//	api.sortBy("SizeMB", reverse)
//}

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

//func (api *IPAddressAPI) Create(value *sacloud.IPAddress) (*sacloud.IPAddress, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.create(api.createRequest(value), res)
//	})
//}

//func (api *IPAddressAPI) Read(id int64) (*sacloud.IPAddress, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.read(id, nil, res)
//	})
//}

//func (api *IPAddressAPI) Update(id int64, value *sacloud.IPAddress) (*sacloud.IPAddress, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.update(id, api.createRequest(value), res)
//	})
//}
//
//func (api *IPAddressAPI) Delete(id int64) (*sacloud.IPAddress, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.delete(id, nil, res)
//	})
//}
//
//func (api *IPAddressAPI) New() *sacloud.IPAddress {
//	return &sacloud.IPAddress{
//	}
//}

/************************************************
  Inner functions
************************************************/

func (api *IPAddressAPI) setStateValue(setFunc func(*sacloud.Request)) *IPAddressAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *IPAddressAPI) request(f func(*sacloud.Response) error) (*sacloud.IPAddress, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.IPAddress, nil
}

func (api *IPAddressAPI) createRequest(value *sacloud.IPAddress) *sacloud.Request {
	req := &sacloud.Request{}
	req.IPAddress = value
	return req
}
