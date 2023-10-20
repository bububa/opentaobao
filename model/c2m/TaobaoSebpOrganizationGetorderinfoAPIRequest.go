package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosebporganizationgetorderinfoAPIRequest 淘小铺机构订单信息 API请求
// taobao.sebp.organization.getorderinfo
//
// 淘小铺合作机构获取机构订单信息，用于机构结算使用
type TaobaosebporganizationgetorderinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 查询实时数据时，必传，开始时间结束时间间隔不能超过4个小时
	_endTime string
	// 查询实时数据时，必传，开始时间不能早于2天前
	_startTime string
	// 第几页
	_pageNum int64
}

// NewTaobaosebporganizationgetorderinfoRequest 初始化TaobaosebporganizationgetorderinfoAPIRequest对象
func NewTaobaosebporganizationgetorderinfoRequest() *TaobaosebporganizationgetorderinfoAPIRequest {
	return &TaobaosebporganizationgetorderinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.organization.getorderinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyDate is ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaosebporganizationgetorderinfoAPIRequest) SetModifyDate(_modifyDate string) error {
	r._modifyDate = _modifyDate
	r.Set("modify_date", _modifyDate)
	return nil
}

// GetModifyDate ModifyDate Getter
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetModifyDate() string {
	return r._modifyDate
}

// SetEndTime is EndTime Setter
// 查询实时数据时，必传，开始时间结束时间间隔不能超过4个小时
func (r *TaobaosebporganizationgetorderinfoAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 查询实时数据时，必传，开始时间不能早于2天前
func (r *TaobaosebporganizationgetorderinfoAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageNum is PageNum Setter
// 第几页
func (r *TaobaosebporganizationgetorderinfoAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaosebporganizationgetorderinfoAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
