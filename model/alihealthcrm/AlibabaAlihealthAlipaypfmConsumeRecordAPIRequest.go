package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest
记录用户每日消耗卡路里总量 API请求
alibaba.alihealth.alipaypfm.consume.record

记录用户每日消耗卡路里总量 */
type AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest struct {
	model.Params
	// 用户健康ID
	_userId int64
	// 用户消耗卡路里总量
	_energy int64
	// 记录日期, 格式: yyyy-MM-dd
	_date string
}

// NewAlibabaAlihealthAlipaypfmConsumeRecordRequest 初始化AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest对象
func NewAlibabaAlihealthAlipaypfmConsumeRecordRequest() *AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest {
	return &AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alipaypfm.consume.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 用户健康ID
func (r *AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is Energy Setter
// 用户消耗卡路里总量
func (r *AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) SetEnergy(_energy int64) error {
	r._energy = _energy
	r.Set("energy", _energy)
	return nil
}

// Get Energy Getter
func (r AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) GetEnergy() int64 {
	return r._energy
}

// Set is Date Setter
// 记录日期, 格式: yyyy-MM-dd
func (r *AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// Get Date Getter
func (r AlibabaAlihealthAlipaypfmConsumeRecordAPIRequest) GetDate() string {
	return r._date
}
