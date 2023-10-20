package film

import (
	"sync"
)

// ResultListModel 结构体
type ResultListModel struct {
	// 返回值
	AccountList []TaobaoFilmAccountPhoneQueryModel `json:"account_list,omitempty" xml:"account_list>taobao_film_account_phone_query_model,omitempty"`
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 忽略
	ReturnUrl string `json:"return_url,omitempty" xml:"return_url,omitempty"`
}

var poolResultListModel = sync.Pool{
	New: func() any {
		return new(ResultListModel)
	},
}

// GetResultListModel() 从对象池中获取ResultListModel
func GetResultListModel() *ResultListModel {
	return poolResultListModel.Get().(*ResultListModel)
}

// ReleaseResultListModel 释放ResultListModel
func ReleaseResultListModel(v *ResultListModel) {
	v.AccountList = v.AccountList[:0]
	v.ReturnCode = ""
	v.RequestId = ""
	v.ReturnMessage = ""
	v.ReturnUrl = ""
	poolResultListModel.Put(v)
}
