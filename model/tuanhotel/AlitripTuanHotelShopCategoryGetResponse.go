package tuanhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家店铺类目查询 APIResponse
alitrip.tuan.hotel.shop.category.get

查询商家店铺类目信息
*/
type AlitripTuanHotelShopCategoryGetAPIResponse struct {
    model.CommonResponse
    AlitripTuanHotelShopCategoryGetResponse
}

type AlitripTuanHotelShopCategoryGetResponse struct {
    XMLName xml.Name `xml:"alitrip_tuan_hotel_shop_category_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 一级类目列表
    
    TopRootShopCategoryList   []TopRootShopCategoryVoList `json:"top_root_shop_category_list,omitempty" xml:"top_root_shop_category_list>top_root_shop_category_vo_list,omitempty"`
    
    
    // code
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 是否成功
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
