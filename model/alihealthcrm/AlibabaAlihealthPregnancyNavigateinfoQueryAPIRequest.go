package alihealthcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest 查询底导数据 API请求
// alibaba.alihealth.pregnancy.navigateinfo.query
//
// 备孕管理--获取底部导航信息
type AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
}

// NewAlibabaAlihealthPregnancyNavigateinfoQueryRequest 初始化AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest对象
func NewAlibabaAlihealthPregnancyNavigateinfoQueryRequest() *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest {
	return &AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) Reset() {
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.navigateinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPregnancyNavigateinfoQueryRequest()
	},
}

// GetAlibabaAlihealthPregnancyNavigateinfoQueryRequest 从 sync.Pool 获取 AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest
func GetAlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest() *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest {
	return poolAlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest.Get().(*AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest)
}

// ReleaseAlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest 将 AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest(v *AlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPregnancyNavigateinfoQueryAPIRequest.Put(v)
}
