package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTagItemsQueryAPIRequest 打标结果查询-标维度 API请求
// taobao.qimen.tag.items.query
//
// 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoQimenTagItemsQueryAPIRequest struct {
	model.Params
	// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
	_tagType string
	// 备注，string（500）
	_remark string
}

// NewTaobaoQimenTagItemsQueryRequest 初始化TaobaoQimenTagItemsQueryAPIRequest对象
func NewTaobaoQimenTagItemsQueryRequest() *TaobaoQimenTagItemsQueryAPIRequest {
	return &TaobaoQimenTagItemsQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenTagItemsQueryAPIRequest) Reset() {
	r._tagType = ""
	r._remark = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTagItemsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.tag.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTagItemsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenTagItemsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagType is TagType Setter
// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
func (r *TaobaoQimenTagItemsQueryAPIRequest) SetTagType(_tagType string) error {
	r._tagType = _tagType
	r.Set("tag_type", _tagType)
	return nil
}

// GetTagType TagType Getter
func (r TaobaoQimenTagItemsQueryAPIRequest) GetTagType() string {
	return r._tagType
}

// SetRemark is Remark Setter
// 备注，string（500）
func (r *TaobaoQimenTagItemsQueryAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoQimenTagItemsQueryAPIRequest) GetRemark() string {
	return r._remark
}

var poolTaobaoQimenTagItemsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenTagItemsQueryRequest()
	},
}

// GetTaobaoQimenTagItemsQueryRequest 从 sync.Pool 获取 TaobaoQimenTagItemsQueryAPIRequest
func GetTaobaoQimenTagItemsQueryAPIRequest() *TaobaoQimenTagItemsQueryAPIRequest {
	return poolTaobaoQimenTagItemsQueryAPIRequest.Get().(*TaobaoQimenTagItemsQueryAPIRequest)
}

// ReleaseTaobaoQimenTagItemsQueryAPIRequest 将 TaobaoQimenTagItemsQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenTagItemsQueryAPIRequest(v *TaobaoQimenTagItemsQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenTagItemsQueryAPIRequest.Put(v)
}
