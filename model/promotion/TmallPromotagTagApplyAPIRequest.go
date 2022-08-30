package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTagApplyAPIRequest 优惠标签申请 API请求
// tmall.promotag.tag.apply
//
// 创建优惠标签
type TmallPromotagTagApplyAPIRequest struct {
	model.Params
	// 标签结束时间
	_endTime string
	// 标签开始时间
	_startTime string
	// 标签名称。注意在UMP中使用新人群标签top变成大写的“NEW_” 如：老标签是top1234，新标签是NEW_1234 。
	_tagName string
	// 标签用途描述
	_tagDesc string
}

// NewTmallPromotagTagApplyRequest 初始化TmallPromotagTagApplyAPIRequest对象
func NewTmallPromotagTagApplyRequest() *TmallPromotagTagApplyAPIRequest {
	return &TmallPromotagTagApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTagApplyAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.tag.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTagApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEndTime is EndTime Setter
// 标签结束时间
func (r *TmallPromotagTagApplyAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TmallPromotagTagApplyAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 标签开始时间
func (r *TmallPromotagTagApplyAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TmallPromotagTagApplyAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetTagName is TagName Setter
// 标签名称。注意在UMP中使用新人群标签top变成大写的“NEW_” 如：老标签是top1234，新标签是NEW_1234 。
func (r *TmallPromotagTagApplyAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r TmallPromotagTagApplyAPIRequest) GetTagName() string {
	return r._tagName
}

// SetTagDesc is TagDesc Setter
// 标签用途描述
func (r *TmallPromotagTagApplyAPIRequest) SetTagDesc(_tagDesc string) error {
	r._tagDesc = _tagDesc
	r.Set("tag_desc", _tagDesc)
	return nil
}

// GetTagDesc TagDesc Getter
func (r TmallPromotagTagApplyAPIRequest) GetTagDesc() string {
	return r._tagDesc
}
