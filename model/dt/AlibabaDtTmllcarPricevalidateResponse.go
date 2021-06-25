package dt

import (
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
    Response *AlibabaDtTmllcarPricevalidateResponse `json:"alibaba_dt_tmllcar_pricevalidate_response,omitempty"`
}

type AlibabaDtTmllcarPricevalidateResponse struct {

    // result
    Result   *Results `json:"result,omitempty"`

}
