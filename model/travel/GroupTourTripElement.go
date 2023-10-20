package travel

import (
	"sync"
)

// GroupTourTripElement 结构体
type GroupTourTripElement struct {
	// 具体的行程信息，根据type字段，将对象序列化成json串，以字符串的形式赋值给json_str，传到后端,每一个json_str都只能是对应单个对象，不能对应数组  当type=1时： json_str = {     &#34;type&#34;:1, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     &#34;hotelType”:1,//住宿方式，1:酒店/客栈 2:住在交通工具上 3:住宿自理 4:露营     &#34;hotelStarType&#34;:1, //1:酒店星级 2:酒店档次     &#34;hotelStar&#34;:&#34;三星级&#34;, //如果hotelStarType =1:酒店星级，hotelStar取值范围（一星级，二星级，三星级，四星级，五星级）；如果hotelStarType =1:酒店星级，hotelStar取值范围（舒适,高档,豪华,经济连锁，二星及以下）     &#34;hotelCityName&#34;:&#34;北京市&#34;, //酒店所在城市名称     &#34;hotelName&#34;:&#34;如家快捷北京北太平庄店”,//酒店名称     &#34;roomType&#34;:&#34;大床房”, //房型     &#34;tripContentDetails&#34;:{ //该字段选填，imageList为数组类型，住宿图片，desc为住宿说明                           &#34;imageList&#34;:[                                        &#34;https://img.daily.taobaocdn.net/imgextra/i2/2024098454/TB2.BJ9XEw7LKJjyzdKXXaShXXa_!!2024098454.jpg&#34;,                                        &#34;https://img.daily.taobaocdn.net/imgextra/i1/2024098454/TB2Ui4yXEw7LKJjyzdKXXaShXXa_!!2024098454.jpg&#34;                                       ],                           &#34;desc”:”住宿说明”                          } } 当type=2时： json_str = {     &#34;type”:2, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     &#34;activityHour&#34;:10, //活动时间-小时     &#34;activityMinute&#34;:30,//活动时间-分钟,这里是10个小时30分钟     &#34;scenicName&#34;:&#34;八达岭长城”,//景点名称     &#34;scenicCity&#34;:&#34;北京市” //景点所在城市,”classicScenic”:true, &#34;tripContentDetails&#34;:{ //该字段选填，imageList为数组类型，景点图片，desc为景点详细说明                           &#34;imageList&#34;:[                                        &#34;https://img.daily.taobaocdn.net/imgextra/i2/2024098454/TB2.BJ9XEw7LKJjyzdKXXaShXXa_!!2024098454.jpg&#34;,                                        &#34;https://img.daily.taobaocdn.net/imgextra/i1/2024098454/TB2Ui4yXEw7LKJjyzdKXXaShXXa_!!2024098454.jpg&#34;                                       ],                           &#34;desc”:”景点详情”                          } } 当type=3时： json_str ={     &#34;type&#34;:3, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     &#34;foodInclude&#34;:false,//true：包含餐饮，false:不包含餐饮信息     &#34;specialIllustrate&#34;:&#34; 餐饮说明”,//餐饮说明     &#34;foodType&#34;:[ //1:早餐，2:中餐，3:晚餐         1,         2,         3     ] } 当type=4时： json_str ={     &#34;type&#34;:4, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息     &#34;activityHour&#34;:1,//活动时间-小时     &#34;activityMinute&#34;:5,//活动时间-分钟,这里是1个小时5分钟     &#34;shoppingPlace&#34;:&#34;家乐福”,//购物店名称     &#34;shoppingProduct&#34;:&#34;啥都有”//营业产品 }。当type=5时： json_str ={ &#34;type&#34;:5, //包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息，5:自由活动 ,&#34;activityHour&#34;:1,//活动时间-小时 &#34;activityMinute&#34;:5,//活动时间-分钟,这里是1个小时5分钟 &#34;scenicCity&#34;:&#34;杭州”,//活动城市, &#34;activityContent&#34;:&#34;啥都有”//活动推荐 }
	JsonStr string `json:"json_str,omitempty" xml:"json_str,omitempty"`
	// 必填，第几天的行程信息
	Day int64 `json:"day,omitempty" xml:"day,omitempty"`
	// 必填，包含元素类型，1:住宿信息，2:景点，3:餐饮信息，4:购物信息，5:自由活动
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolGroupTourTripElement = sync.Pool{
	New: func() any {
		return new(GroupTourTripElement)
	},
}

// GetGroupTourTripElement() 从对象池中获取GroupTourTripElement
func GetGroupTourTripElement() *GroupTourTripElement {
	return poolGroupTourTripElement.Get().(*GroupTourTripElement)
}

// ReleaseGroupTourTripElement 释放GroupTourTripElement
func ReleaseGroupTourTripElement(v *GroupTourTripElement) {
	v.JsonStr = ""
	v.Day = 0
	v.Type = 0
	poolGroupTourTripElement.Put(v)
}
