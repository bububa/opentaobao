package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpdetaillistaddAPIRequest 营销详情添加 API请求
// taobao.ump.detail.list.add
//
// 批量添加营销活动。替代单条添加营销详情的的API。此接口适用针对某个营销活动的多档设置，会按顺序插入detail。若在整个事务过程中出现断点，会将已插入完成的detail_id返回，注意记录这些id，并将其删除，会对交易过程中的优惠产生影响。
type TaobaoumpdetaillistaddAPIRequest struct {
	model.Params
	// 营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。
	_details string
	// 营销活动id。
	_actId int64
}

// NewTaobaoumpdetaillistaddRequest 初始化TaobaoumpdetaillistaddAPIRequest对象
func NewTaobaoumpdetaillistaddRequest() *TaobaoumpdetaillistaddAPIRequest {
	return &TaobaoumpdetaillistaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpdetaillistaddAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.list.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpdetaillistaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpdetaillistaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetails is Details Setter
// 营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。
func (r *TaobaoumpdetaillistaddAPIRequest) SetDetails(_details string) error {
	r._details = _details
	r.Set("details", _details)
	return nil
}

// GetDetails Details Getter
func (r TaobaoumpdetaillistaddAPIRequest) GetDetails() string {
	return r._details
}

// SetActId is ActId Setter
// 营销活动id。
func (r *TaobaoumpdetaillistaddAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumpdetaillistaddAPIRequest) GetActId() int64 {
	return r._actId
}
