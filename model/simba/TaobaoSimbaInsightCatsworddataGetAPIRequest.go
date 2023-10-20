package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbainsightcatsworddatagetAPIRequest 获取类目下关键词的数据 API请求
// taobao.simba.insight.catsworddata.get
//
// 获取给定词在给定类目下的详细数据
type TaobaosimbainsightcatsworddatagetAPIRequest struct {
	model.Params
	// 需要查询的关键词列表，最大长度100。
	_bidwordList []string
	// 类目id
	_catId string
	// 开始时间，格式只能为：yyyy-MM-dd
	_startDate string
	// 结束时间，格式只能为：yyyy-MM-dd
	_endDate string
}

// NewTaobaosimbainsightcatsworddatagetRequest 初始化TaobaosimbainsightcatsworddatagetAPIRequest对象
func NewTaobaosimbainsightcatsworddatagetRequest() *TaobaosimbainsightcatsworddatagetAPIRequest {
	return &TaobaosimbainsightcatsworddatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.insight.catsworddata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordList is BidwordList Setter
// 需要查询的关键词列表，最大长度100。
func (r *TaobaosimbainsightcatsworddatagetAPIRequest) SetBidwordList(_bidwordList []string) error {
	r._bidwordList = _bidwordList
	r.Set("bidword_list", _bidwordList)
	return nil
}

// GetBidwordList BidwordList Getter
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetBidwordList() []string {
	return r._bidwordList
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaosimbainsightcatsworddatagetAPIRequest) SetCatId(_catId string) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetCatId() string {
	return r._catId
}

// SetStartDate is StartDate Setter
// 开始时间，格式只能为：yyyy-MM-dd
func (r *TaobaosimbainsightcatsworddatagetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间，格式只能为：yyyy-MM-dd
func (r *TaobaosimbainsightcatsworddatagetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaosimbainsightcatsworddatagetAPIRequest) GetEndDate() string {
	return r._endDate
}
