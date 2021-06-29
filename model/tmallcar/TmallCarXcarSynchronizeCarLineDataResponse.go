package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的爱卡车型配置数据 APIResponse
tmall.car.xcar.synchronize.car.line.data

同步我的爱卡车系配置数据到天猫汽车
*/
type TmallCarXcarSynchronizeCarLineDataAPIResponse struct {
    model.CommonResponse
    TmallCarXcarSynchronizeCarLineDataResponse
}

type TmallCarXcarSynchronizeCarLineDataResponse struct {
    XMLName xml.Name `xml:"tmall_car_xcar_synchronize_car_line_data_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象总体信息
    
    Result   *ResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
