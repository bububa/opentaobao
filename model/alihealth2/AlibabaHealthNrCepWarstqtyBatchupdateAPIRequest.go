package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest 批量更新ISV库存 API请求
// alibaba.health.nr.cep.warstqty.batchupdate
//
// 青岛医保服务-ISV批量更新孔雀翎中库存数据
type AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest struct {
	model.Params
	// 库存更新对象
	_warStqtyList []TopIsvStqtyLstDto
}

// NewAlibabaHealthNrCepWarstqtyBatchupdateRequest 初始化AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest对象
func NewAlibabaHealthNrCepWarstqtyBatchupdateRequest() *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest {
	return &AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) Reset() {
	r._warStqtyList = r._warStqtyList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.cep.warstqty.batchupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarStqtyList is WarStqtyList Setter
// 库存更新对象
func (r *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) SetWarStqtyList(_warStqtyList []TopIsvStqtyLstDto) error {
	r._warStqtyList = _warStqtyList
	r.Set("war_stqty_list", _warStqtyList)
	return nil
}

// GetWarStqtyList WarStqtyList Getter
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetWarStqtyList() []TopIsvStqtyLstDto {
	return r._warStqtyList
}

var poolAlibabaHealthNrCepWarstqtyBatchupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthNrCepWarstqtyBatchupdateRequest()
	},
}

// GetAlibabaHealthNrCepWarstqtyBatchupdateRequest 从 sync.Pool 获取 AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest
func GetAlibabaHealthNrCepWarstqtyBatchupdateAPIRequest() *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest {
	return poolAlibabaHealthNrCepWarstqtyBatchupdateAPIRequest.Get().(*AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest)
}

// ReleaseAlibabaHealthNrCepWarstqtyBatchupdateAPIRequest 将 AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthNrCepWarstqtyBatchupdateAPIRequest(v *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) {
	v.Reset()
	poolAlibabaHealthNrCepWarstqtyBatchupdateAPIRequest.Put(v)
}
