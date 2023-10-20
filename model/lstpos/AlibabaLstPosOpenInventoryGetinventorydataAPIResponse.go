package lstpos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenInventoryGetinventorydataAPIResponse 商品库存只读接口(最多20条库存信息) API返回值
// alibaba.lst.pos.open.inventory.getinventorydata
//
// 商品库存只读接口(最多20条库存信息)
type AlibabaLstPosOpenInventoryGetinventorydataAPIResponse struct {
	model.CommonResponse
	AlibabaLstPosOpenInventoryGetinventorydataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenInventoryGetinventorydataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstPosOpenInventoryGetinventorydataAPIResponseModel).Reset()
}

// AlibabaLstPosOpenInventoryGetinventorydataAPIResponseModel is 商品库存只读接口(最多20条库存信息) 成功返回结果
type AlibabaLstPosOpenInventoryGetinventorydataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_inventory_getinventorydata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果对象
	Result *AlibabaLstPosOpenInventoryGetinventorydataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenInventoryGetinventorydataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstPosOpenInventoryGetinventorydataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenInventoryGetinventorydataAPIResponse)
	},
}

// GetAlibabaLstPosOpenInventoryGetinventorydataAPIResponse 从 sync.Pool 获取 AlibabaLstPosOpenInventoryGetinventorydataAPIResponse
func GetAlibabaLstPosOpenInventoryGetinventorydataAPIResponse() *AlibabaLstPosOpenInventoryGetinventorydataAPIResponse {
	return poolAlibabaLstPosOpenInventoryGetinventorydataAPIResponse.Get().(*AlibabaLstPosOpenInventoryGetinventorydataAPIResponse)
}

// ReleaseAlibabaLstPosOpenInventoryGetinventorydataAPIResponse 将 AlibabaLstPosOpenInventoryGetinventorydataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstPosOpenInventoryGetinventorydataAPIResponse(v *AlibabaLstPosOpenInventoryGetinventorydataAPIResponse) {
	v.Reset()
	poolAlibabaLstPosOpenInventoryGetinventorydataAPIResponse.Put(v)
}
