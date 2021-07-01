package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
除度假线路、门票以外的其他类目商品维护接口（商品ID重复将自动更新） 
alitrip.travel.gereralitem.update

除度假线路、门票以外的商品维护接口；目前该接口支持以下类目；
（123740001：客栈周边交通服务、125038002：旅行设备/GPS/相机租赁、50018298：船票、124084006：酒店商品升级差价、125228016：预约卡券、50011954：旅游服务、50012913：酒店优惠券、50214003：旅游会员卡/酒店会员卡、50012917：巴士/地铁/交通卡/一卡通、50134002：代客烧香/还愿、50026091：境外火车票、123742001：客栈周边餐饮服务、50019817：海外套餐（该类目已废弃）、125210016：团建/outing、124212017：其他预定、50025888：机场行李托运取送寄存、50025831：旅游年票/年卡、124142009：旅游商品升级差价/押金、123744001：客栈周边其他服务、50012762：广深口岸港澳送关服务、50025880：旅行拍照/婚纱摄影、123166001：酒店餐饮美食（该类目已废弃）、50668002：手绘地图/明信片、50024210：旅游购物/纪念品、50024208：酒店用品、50024215：购物折扣卡券、50025878：酒店SPA/足浴/温泉、50024212：旅游必备品、123738001：客栈周边票务服务、123164002：游泳健身（该类目已废弃）、50686003：机票增值产品、123164001：酒店SPA（该类目已废弃）、124832008：美食卡券/酒店餐饮卡券、125408001：旅游定制服务、50018112：旅行社/网站优惠券、124258004：酒店客房优惠券（该类目已废弃）、50104001：机场周边停车位、124730002：内机机场/火车站送关服务、124090010：其他）
*/
func AlitripTravelGereralitemUpdate(clt *core.SDKClient, req *travel.AlitripTravelGereralitemUpdateAPIRequest, session string) (*travel.AlitripTravelGereralitemUpdateAPIResponse, error) {
    var resp travel.AlitripTravelGereralitemUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
