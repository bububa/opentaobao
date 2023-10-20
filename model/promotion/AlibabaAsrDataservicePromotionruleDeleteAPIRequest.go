package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaasrdataservicepromotionruledeleteAPIRequest 优惠规则删除 API请求
// alibaba.asr.dataservice.promotionrule.delete
//
// 删除优惠规则，例如星巴克删除优惠规则
type AlibabaasrdataservicepromotionruledeleteAPIRequest struct {
	model.Params
	// poskey
	_posKey int64
}

// NewAlibabaasrdataservicepromotionruledeleteRequest 初始化AlibabaasrdataservicepromotionruledeleteAPIRequest对象
func NewAlibabaasrdataservicepromotionruledeleteRequest() *AlibabaasrdataservicepromotionruledeleteAPIRequest {
	return &AlibabaasrdataservicepromotionruledeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaasrdataservicepromotionruledeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaasrdataservicepromotionruledeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaasrdataservicepromotionruledeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosKey is PosKey Setter
// poskey
func (r *AlibabaasrdataservicepromotionruledeleteAPIRequest) SetPosKey(_posKey int64) error {
	r._posKey = _posKey
	r.Set("pos_key", _posKey)
	return nil
}

// GetPosKey PosKey Getter
func (r AlibabaasrdataservicepromotionruledeleteAPIRequest) GetPosKey() int64 {
	return r._posKey
}
