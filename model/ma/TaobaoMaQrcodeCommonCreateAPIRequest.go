package ma

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMaQrcodeCommonCreateAPIRequest
创建码平台常用二维码 API请求
taobao.ma.qrcode.common.create

创建码平台对外提供的常用二维码接口，不适于码平台业务类型的码创建，如不支持包裹码、媒体码等，业务类型的码需要单独提供API。 */
type TaobaoMaQrcodeCommonCreateAPIRequest struct {
	model.Params
	// 要制作的二维码业务类型：page:无线页面类型item:宝贝ID类型url:普通的URL连接类型shop:店铺ID类型
	_type string
	// 二维码的内容之一，由type决定：type=page时，content传入无线页面的URL连接内容；type=item时，content传入宝贝数字ID；type=url时，content传入普通的URL连接内容；type=shop时，content传入店铺ID；
	_content string
	// 二维码名字，即创建的二维码，在码平台上显示记录的名字。
	_name string
	// 二维码需要布点的位置，方便用户在码平台上可以区分看到不同布点的扫码数据情况；列表值，用半角','号分割，单个渠道名不能超过16字符。
	_channelName string
	// 二维码的样式名，支持普通码的颜色或官方模板的模板名；普通码的颜色可选输入：“000000”(黑色)、“EF4F2B”(橙色);官方模板的可选输入（实际尺寸比样例大）：“ww_color.png“ 尺寸290x320，样例：http://gtms03.alicdn.com/tps/i3/T1YLPRFRXXXXbsbYwb-100-102.png；“tb_scan.png“ 尺寸290x320，样例：http://gtms01.alicdn.com/tps/i1/T14vsEFThdXXbsbYwb-100-102.png；“ww_hide_color.png“  尺寸200x263，样例：http://gtms04.alicdn.com/tps/i4/TB1URvlFVXXXXbRXFXXwxcf6pXX-76-100.png；“tmall_hide_color.png“ 尺寸200x263，样例：http://gtms01.alicdn.com/tps/i1/TB1S5PiFVXXXXacXVXXwxcf6pXX-76-100.png。
	_style string
	// 二维码尺寸，只支持普通二维码，不支持官方模板，单位为像素，最小为60×60，最大为300×300，建议175×175。官方模板大小尺寸见style说明。
	_size int64
	// 是否需要矢量图，如果需要矢量图，设置为true；只支持普通二维码，官方模板不支持矢量图
	_needEps bool
	// 二维码的logo地址，只允许淘宝官方图片空间的图片地址，其他非图片空间图片不支持。官方淘logo图片地址：http://img01.taobaocdn.com/imgextra/T1Od8YFT8eXXXXXXXX。
	_logo string
}

// New
