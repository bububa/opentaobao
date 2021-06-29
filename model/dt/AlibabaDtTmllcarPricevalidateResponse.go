package dt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线索报价价格校验 APIResponse
alibaba.dt.tmllcar.pricevalidate

根据选定的车型和城市，校验汽车价格是否通过
入参：车型ID，城市名称，价格
输出：N 校验失败，校验成功不返回值
*/
type AlibabaDtTmllcarPricevalidateAPIResponse struct {
    model.CommonResponse
    AlibabaDtTmllcarPricevalidateResponse
}

type AlibabaDtTmllcarPricevalidateResponse struct {
    XMLName xml.Name `xml:"alibaba_dt_tmllcar_pricevalidate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaDtTmllcarPricevalidateResults `json:"result,omitempty" xml:"result,omitempty"`

    
}
