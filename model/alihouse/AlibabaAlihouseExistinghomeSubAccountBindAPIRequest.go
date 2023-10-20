package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeSubAccountBindAPIRequest 子账号入驻 API请求
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
type AlibabaAlihouseExistinghomeSubAccountBindAPIRequest struct {
	model.Params
	// dto
	_subAccountReqDto *SubAccountReqDto
}

// NewAlibabaAlihouseExistinghomeSubAccountBindRequest 初始化AlibabaAlihouseExistinghomeSubAccountBindAPIRequest对象
func NewAlibabaAlihouseExistinghomeSubAccountBindRequest() *AlibabaAlihouseExistinghomeSubAccountBindAPIRequest {
	return &AlibabaAlihouseExistinghomeSubAccountBindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) Reset() {
	r._subAccountReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.sub.account.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubAccountReqDto is SubAccountReqDto Setter
// dto
func (r *AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) SetSubAccountReqDto(_subAccountReqDto *SubAccountReqDto) error {
	r._subAccountReqDto = _subAccountReqDto
	r.Set("sub_account_req_dto", _subAccountReqDto)
	return nil
}

// GetSubAccountReqDto SubAccountReqDto Getter
func (r AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) GetSubAccountReqDto() *SubAccountReqDto {
	return r._subAccountReqDto
}

var poolAlibabaAlihouseExistinghomeSubAccountBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeSubAccountBindRequest()
	},
}

// GetAlibabaAlihouseExistinghomeSubAccountBindRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeSubAccountBindAPIRequest
func GetAlibabaAlihouseExistinghomeSubAccountBindAPIRequest() *AlibabaAlihouseExistinghomeSubAccountBindAPIRequest {
	return poolAlibabaAlihouseExistinghomeSubAccountBindAPIRequest.Get().(*AlibabaAlihouseExistinghomeSubAccountBindAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeSubAccountBindAPIRequest 将 AlibabaAlihouseExistinghomeSubAccountBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeSubAccountBindAPIRequest(v *AlibabaAlihouseExistinghomeSubAccountBindAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeSubAccountBindAPIRequest.Put(v)
}
