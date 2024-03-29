package c2m

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSebpOrganizationGetorderinfoAPIRequest 淘小铺机构订单信息 API请求
// taobao.sebp.organization.getorderinfo
//
// 淘小铺合作机构获取机构订单信息，用于机构结算使用
type TaobaoSebpOrganizationGetorderinfoAPIRequest struct {
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

// NewTaobaoSebpOrganizationGetorderinfoRequest 初始化TaobaoSebpOrganizationGetorderinfoAPIRequest对象
func NewTaobaoSebpOrganizationGetorderinfoRequest() *TaobaoSebpOrganizationGetorderinfoAPIRequest {
	return &TaobaoSebpOrganizationGetorderinfoAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSebpOrganizationGetorderinfoAPIRequest) Reset() {
	r._modifyDate = ""
	r._endTime = ""
	r._startTime = ""
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.organization.getorderinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyDate is ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaoSebpOrganizationGetorderinfoAPIRequest) SetModifyDate(_modifyDate string) error {
	r._modifyDate = _modifyDate
	r.Set("modify_date", _modifyDate)
	return nil
}

// GetModifyDate ModifyDate Getter
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetModifyDate() string {
	return r._modifyDate
}

// SetEndTime is EndTime Setter
// 查询实时数据时，必传，开始时间结束时间间隔不能超过4个小时
func (r *TaobaoSebpOrganizationGetorderinfoAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 查询实时数据时，必传，开始时间不能早于2天前
func (r *TaobaoSebpOrganizationGetorderinfoAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageNum is PageNum Setter
// 第几页
func (r *TaobaoSebpOrganizationGetorderinfoAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoSebpOrganizationGetorderinfoAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolTaobaoSebpOrganizationGetorderinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSebpOrganizationGetorderinfoRequest()
	},
}

// GetTaobaoSebpOrganizationGetorderinfoRequest 从 sync.Pool 获取 TaobaoSebpOrganizationGetorderinfoAPIRequest
func GetTaobaoSebpOrganizationGetorderinfoAPIRequest() *TaobaoSebpOrganizationGetorderinfoAPIRequest {
	return poolTaobaoSebpOrganizationGetorderinfoAPIRequest.Get().(*TaobaoSebpOrganizationGetorderinfoAPIRequest)
}

// ReleaseTaobaoSebpOrganizationGetorderinfoAPIRequest 将 TaobaoSebpOrganizationGetorderinfoAPIRequest 放入 sync.Pool
func ReleaseTaobaoSebpOrganizationGetorderinfoAPIRequest(v *TaobaoSebpOrganizationGetorderinfoAPIRequest) {
	v.Reset()
	poolTaobaoSebpOrganizationGetorderinfoAPIRequest.Put(v)
}
