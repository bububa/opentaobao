package shenjing

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest pad获取二维码 API请求
// alibaba.ib.shenjing.visitor.pad.getqrcodelink
//
// pad获取二维码链接。扫码录入人脸。
type AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest struct {
	model.Params
	// 终端id
	_termId string
}

// NewAlibabaIbShenjingVisitorPadGetqrcodelinkRequest 初始化AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest对象
func NewAlibabaIbShenjingVisitorPadGetqrcodelinkRequest() *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest {
	return &AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) Reset() {
	r._termId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.getqrcodelink"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTermId is TermId Setter
// 终端id
func (r *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) SetTermId(_termId string) error {
	r._termId = _termId
	r.Set("term_id", _termId)
	return nil
}

// GetTermId TermId Getter
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetTermId() string {
	return r._termId
}

var poolAlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIbShenjingVisitorPadGetqrcodelinkRequest()
	},
}

// GetAlibabaIbShenjingVisitorPadGetqrcodelinkRequest 从 sync.Pool 获取 AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest
func GetAlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest() *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest {
	return poolAlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest.Get().(*AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest)
}

// ReleaseAlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest 将 AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest 放入 sync.Pool
func ReleaseAlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest(v *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) {
	v.Reset()
	poolAlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest.Put(v)
}
