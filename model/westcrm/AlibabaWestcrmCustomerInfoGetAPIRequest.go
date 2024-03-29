package westcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmCustomerInfoGetAPIRequest 会员信息查询接口 API请求
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
type AlibabaWestcrmCustomerInfoGetAPIRequest struct {
	model.Params
	// 支付宝id
	_alipayId string
	// 园区id
	_campusId int64
	// 会员id
	_ibUserId int64
}

// NewAlibabaWestcrmCustomerInfoGetRequest 初始化AlibabaWestcrmCustomerInfoGetAPIRequest对象
func NewAlibabaWestcrmCustomerInfoGetRequest() *AlibabaWestcrmCustomerInfoGetAPIRequest {
	return &AlibabaWestcrmCustomerInfoGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWestcrmCustomerInfoGetAPIRequest) Reset() {
	r._alipayId = ""
	r._campusId = 0
	r._ibUserId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmCustomerInfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.customer.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmCustomerInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWestcrmCustomerInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayId is AlipayId Setter
// 支付宝id
func (r *AlibabaWestcrmCustomerInfoGetAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// GetAlipayId AlipayId Getter
func (r AlibabaWestcrmCustomerInfoGetAPIRequest) GetAlipayId() string {
	return r._alipayId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaWestcrmCustomerInfoGetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaWestcrmCustomerInfoGetAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetIbUserId is IbUserId Setter
// 会员id
func (r *AlibabaWestcrmCustomerInfoGetAPIRequest) SetIbUserId(_ibUserId int64) error {
	r._ibUserId = _ibUserId
	r.Set("ib_user_id", _ibUserId)
	return nil
}

// GetIbUserId IbUserId Getter
func (r AlibabaWestcrmCustomerInfoGetAPIRequest) GetIbUserId() int64 {
	return r._ibUserId
}

var poolAlibabaWestcrmCustomerInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWestcrmCustomerInfoGetRequest()
	},
}

// GetAlibabaWestcrmCustomerInfoGetRequest 从 sync.Pool 获取 AlibabaWestcrmCustomerInfoGetAPIRequest
func GetAlibabaWestcrmCustomerInfoGetAPIRequest() *AlibabaWestcrmCustomerInfoGetAPIRequest {
	return poolAlibabaWestcrmCustomerInfoGetAPIRequest.Get().(*AlibabaWestcrmCustomerInfoGetAPIRequest)
}

// ReleaseAlibabaWestcrmCustomerInfoGetAPIRequest 将 AlibabaWestcrmCustomerInfoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWestcrmCustomerInfoGetAPIRequest(v *AlibabaWestcrmCustomerInfoGetAPIRequest) {
	v.Reset()
	poolAlibabaWestcrmCustomerInfoGetAPIRequest.Put(v)
}
