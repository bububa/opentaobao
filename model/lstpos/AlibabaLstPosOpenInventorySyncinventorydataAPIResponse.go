package lstpos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenInventorySyncinventorydataAPIResponse 商品库存修改同步接口(最多20条库存信息) API返回值
// alibaba.lst.pos.open.inventory.syncinventorydata
//
// 商品库存修改同步接口(最多20条库存信息)
type AlibabaLstPosOpenInventorySyncinventorydataAPIResponse struct {
	model.CommonResponse
	AlibabaLstPosOpenInventorySyncinventorydataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenInventorySyncinventorydataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstPosOpenInventorySyncinventorydataAPIResponseModel).Reset()
}

// AlibabaLstPosOpenInventorySyncinventorydataAPIResponseModel is 商品库存修改同步接口(最多20条库存信息) 成功返回结果
type AlibabaLstPosOpenInventorySyncinventorydataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_inventory_syncinventorydata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabaLstPosOpenInventorySyncinventorydataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenInventorySyncinventorydataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstPosOpenInventorySyncinventorydataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenInventorySyncinventorydataAPIResponse)
	},
}

// GetAlibabaLstPosOpenInventorySyncinventorydataAPIResponse 从 sync.Pool 获取 AlibabaLstPosOpenInventorySyncinventorydataAPIResponse
func GetAlibabaLstPosOpenInventorySyncinventorydataAPIResponse() *AlibabaLstPosOpenInventorySyncinventorydataAPIResponse {
	return poolAlibabaLstPosOpenInventorySyncinventorydataAPIResponse.Get().(*AlibabaLstPosOpenInventorySyncinventorydataAPIResponse)
}

// ReleaseAlibabaLstPosOpenInventorySyncinventorydataAPIResponse 将 AlibabaLstPosOpenInventorySyncinventorydataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstPosOpenInventorySyncinventorydataAPIResponse(v *AlibabaLstPosOpenInventorySyncinventorydataAPIResponse) {
	v.Reset()
	poolAlibabaLstPosOpenInventorySyncinventorydataAPIResponse.Put(v)
}
