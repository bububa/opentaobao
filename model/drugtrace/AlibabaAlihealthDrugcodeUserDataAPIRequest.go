package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeUserDataAPIRequest 西安杨森同步用户行为接口 API请求
// alibaba.alihealth.drugcode.user.data
//
// 西安杨森同步用户行为接口
type AlibabaAlihealthDrugcodeUserDataAPIRequest struct {
	model.Params
	// 用户信息
	_list []HaoxinqingDataDto
	// 企业ID，用户区分 appkey下不同企业数据隔离的
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeUserDataRequest 初始化AlibabaAlihealthDrugcodeUserDataAPIRequest对象
func NewAlibabaAlihealthDrugcodeUserDataRequest() *AlibabaAlihealthDrugcodeUserDataAPIRequest {
	return &AlibabaAlihealthDrugcodeUserDataAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeUserDataAPIRequest) Reset() {
	r._list = r._list[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.user.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetList is List Setter
// 用户信息
func (r *AlibabaAlihealthDrugcodeUserDataAPIRequest) SetList(_list []HaoxinqingDataDto) error {
	r._list = _list
	r.Set("list", _list)
	return nil
}

// GetList List Getter
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetList() []HaoxinqingDataDto {
	return r._list
}

// SetRefEntId is RefEntId Setter
// 企业ID，用户区分 appkey下不同企业数据隔离的
func (r *AlibabaAlihealthDrugcodeUserDataAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugcodeUserDataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeUserDataRequest()
	},
}

// GetAlibabaAlihealthDrugcodeUserDataRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeUserDataAPIRequest
func GetAlibabaAlihealthDrugcodeUserDataAPIRequest() *AlibabaAlihealthDrugcodeUserDataAPIRequest {
	return poolAlibabaAlihealthDrugcodeUserDataAPIRequest.Get().(*AlibabaAlihealthDrugcodeUserDataAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeUserDataAPIRequest 将 AlibabaAlihealthDrugcodeUserDataAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeUserDataAPIRequest(v *AlibabaAlihealthDrugcodeUserDataAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeUserDataAPIRequest.Put(v)
}
