package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest 登陆标识信息同步 API请求
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
type AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest struct {
	model.Params
	// 身份信息
	_identityInfo *IdentityInfoDto
}

// NewAlibabaAlihouseExistinghomeIdentifyingSyncRequest 初始化AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeIdentifyingSyncRequest() *AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) Reset() {
	r._identityInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.identifying.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentityInfo is IdentityInfo Setter
// 身份信息
func (r *AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) SetIdentityInfo(_identityInfo *IdentityInfoDto) error {
	r._identityInfo = _identityInfo
	r.Set("identity_info", _identityInfo)
	return nil
}

// GetIdentityInfo IdentityInfo Getter
func (r AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) GetIdentityInfo() *IdentityInfoDto {
	return r._identityInfo
}

var poolAlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeIdentifyingSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeIdentifyingSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest
func GetAlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest() *AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest 将 AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest(v *AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest.Put(v)
}
