package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
影视桌面运营通用接口 API返回值 
yunos.tvpubadmin.epg.desktop.operation

影视桌面运营通用接口
*/
type YunosTvpubadminEpgDesktopOperationAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminEpgDesktopOperationAPIResponseModel
}

// 影视桌面运营通用接口 成功返回结果
type YunosTvpubadminEpgDesktopOperationAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_epg_desktop_operation_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 具体返回结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
