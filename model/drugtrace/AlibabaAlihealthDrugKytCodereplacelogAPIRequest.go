package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodereplacelogAPIRequest 码替换记录查询 API请求
// alibaba.alihealth.drug.kyt.codereplacelog
//
// 码替换记录查询
type AlibabaAlihealthDrugKytCodereplacelogAPIRequest struct {
	model.Params
	// 企业ref_ent_id（码申请企业）
	_refEntId string
	// 开始时间（最大查询区间一年）
	_beginTime string
	// 截至时间（最大查询区间一年）
	_endTime string
	// 追溯码（不区分新码、旧码）
	_code string
	// 药品ID
	_drugEntBaseInfoId string
	// 页数（默认一页20条）
	_page int64
}

// NewAlibabaAlihealthDrugKytCodereplacelogRequest 初始化AlibabaAlihealthDrugKytCodereplacelogAPIRequest对象
func NewAlibabaAlihealthDrugKytCodereplacelogRequest() *AlibabaAlihealthDrugKytCodereplacelogAPIRequest {
	return &AlibabaAlihealthDrugKytCodereplacelogAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) Reset() {
	r._refEntId = ""
	r._beginTime = ""
	r._endTime = ""
	r._code = ""
	r._drugEntBaseInfoId = ""
	r._page = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codereplacelog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id（码申请企业）
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBeginTime is BeginTime Setter
// 开始时间（最大查询区间一年）
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 截至时间（最大查询区间一年）
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCode is Code Setter
// 追溯码（不区分新码、旧码）
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetCode() string {
	return r._code
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetPage is Page Setter
// 页数（默认一页20条）
func (r *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytCodereplacelogAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugKytCodereplacelogAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytCodereplacelogRequest()
	},
}

// GetAlibabaAlihealthDrugKytCodereplacelogRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodereplacelogAPIRequest
func GetAlibabaAlihealthDrugKytCodereplacelogAPIRequest() *AlibabaAlihealthDrugKytCodereplacelogAPIRequest {
	return poolAlibabaAlihealthDrugKytCodereplacelogAPIRequest.Get().(*AlibabaAlihealthDrugKytCodereplacelogAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytCodereplacelogAPIRequest 将 AlibabaAlihealthDrugKytCodereplacelogAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodereplacelogAPIRequest(v *AlibabaAlihealthDrugKytCodereplacelogAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodereplacelogAPIRequest.Put(v)
}
