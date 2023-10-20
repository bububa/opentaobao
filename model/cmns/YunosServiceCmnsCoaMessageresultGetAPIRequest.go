package cmns

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageresultGetAPIRequest CMNS消息发送到达查询 API请求
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
type YunosServiceCmnsCoaMessageresultGetAPIRequest struct {
	model.Params
	// 消息ID
	_mid int64
}

// NewYunosServiceCmnsCoaMessageresultGetRequest 初始化YunosServiceCmnsCoaMessageresultGetAPIRequest对象
func NewYunosServiceCmnsCoaMessageresultGetRequest() *YunosServiceCmnsCoaMessageresultGetAPIRequest {
	return &YunosServiceCmnsCoaMessageresultGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaMessageresultGetAPIRequest) Reset() {
	r._mid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.messageresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMid is Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageresultGetAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosServiceCmnsCoaMessageresultGetAPIRequest) GetMid() int64 {
	return r._mid
}

var poolYunosServiceCmnsCoaMessageresultGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaMessageresultGetRequest()
	},
}

// GetYunosServiceCmnsCoaMessageresultGetRequest 从 sync.Pool 获取 YunosServiceCmnsCoaMessageresultGetAPIRequest
func GetYunosServiceCmnsCoaMessageresultGetAPIRequest() *YunosServiceCmnsCoaMessageresultGetAPIRequest {
	return poolYunosServiceCmnsCoaMessageresultGetAPIRequest.Get().(*YunosServiceCmnsCoaMessageresultGetAPIRequest)
}

// ReleaseYunosServiceCmnsCoaMessageresultGetAPIRequest 将 YunosServiceCmnsCoaMessageresultGetAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaMessageresultGetAPIRequest(v *YunosServiceCmnsCoaMessageresultGetAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaMessageresultGetAPIRequest.Put(v)
}
