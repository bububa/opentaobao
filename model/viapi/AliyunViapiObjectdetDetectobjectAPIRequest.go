package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiObjectdetDetectobjectAPIRequest 物体检测 API请求
// aliyun.viapi.objectdet.detectobject
//
// 检测图像中的物体。90类物体：
// (0：&#39;human&#39;：人体, 1：&#39;sneakers&#39;：胶底运动鞋, 2：&#39;chair&#39;：椅子, 3：&#39;hat&#39;：帽子, 4：&#39;lamp&#39;：灯, 5：&#39;cabinet/shelf&#39;：橱柜/ 架子, 6：&#39;car&#39;：汽车, 7：&#39;glasses&#39;：眼镜, 8：&#39;picture/frame&#39;：照片/图画, 9：&#39;street lights&#39;：街灯, 10：&#39;helmet&#39;：头盔, 11：&#39;pillow&#39;：枕头, 12：&#39;glove&#39;：手套, 13：&#39;potted plant&#39;：盆栽植物, 14：&#39;flower&#39;：花, 15：&#39;monitor&#39;：显示屏, 16：&#39;plants pot/vase&#39;：花盆, 17：&#39;boots&#39;：靴子, 18：&#39;umbrella&#39;：伞, 19：&#39;boat&#39;：小船, 20：&#39;flag&#39;：旗帜, 21：&#39;speaker&#39;：扬声器/话筒, 22：&#39;trash bin/can&#39;：垃圾桶, 23：&#39;backpack&#39;： 双肩背包, 24：&#39;sofa&#39;：沙发, 25：&#39;belt&#39;：腰带, 26：&#39;carpet&#39;：地毯, 27：&#39;coffee table&#39;：咖啡桌/茶几, 28：&#39;tie&#39;： 领带, 29：&#39;bed&#39;： 床, 30：&#39;traffic light&#39;：红绿灯, 31：&#39;necklace&#39;：项链, 32：&#39;mirror&#39;：镜子, 33：&#39;bicycle&#39;：自行车, 34：&#39;watch&#39;：手表, 35：&#39;horse&#39;：马, 36：&#39;traffic sign&#39;：交通标志, 37：&#39;stuffed animal&#39;：填充玩具动物, 38：&#39;motorbike/motorcycle&#39;：摩托车, 39：&#39;wild bird&#39;：鸟, 40：&#39;laptop&#39;：笔记本电脑, 41：&#39;cow&#39;：奶牛, 42：&#39;clock&#39;：时钟, 43：&#39;bus&#39;：公共汽车, 44：&#39;nightstand&#39;：床头柜, 45：&#39;sheep&#39;：绵羊, 46：&#39;traffic cone&#39;：锥形交通路标, 47：&#39;keyboard&#39;：键盘, 48：&#39;hockey stick&#39;：曲棍球球棍, 49：&#39;fan&#39;：电扇, 50：&#39;dog&#39;：狗, 51：&#39;blackboard/whiteboard&#39;：白板/黑板, 52：&#39;mouse&#39;：鼠标, 53：&#39;telephone&#39;：电话, 54：&#39;airplane&#39;：飞机, 55：&#39;skis&#39;：滑雪板, 56：&#39;soccer&#39;：英式足球, 57：&#39;combine with glove&#39;：棒球手套, 58：&#39;train&#39;：火车, 59：&#39;tent&#39;：帐篷, 60：&#39;sailboat&#39;：帆船, 61：&#39;kite&#39;：风筝, 62：&#39;computer box&#39;：计算机主机机箱, 63：&#39;elephant&#39;：大象, 64：&#39;stroller&#39;：折叠式婴儿车, 65：&#39;baseball bat&#39;：棒球棒, 66：&#39;skateboard&#39;：溜冰板, 67：&#39;surfboard&#39;：冲浪板, 68：&#39;cat&#39;：猫, 69：&#39;zebra&#39;：斑马, 70：&#39;sports car&#39;：跑车, 71：&#39;giraffe&#39;：长颈鹿, 72：&#39;radiator&#39;：散热器, 73：&#39;tennis racket&#39;：网球拍, 74：&#39;skating and skiing shoes&#39;：溜冰鞋, 75：&#39;baseball&#39;：棒球, 76：&#39;american football&#39;：美式橄榄球, 77：&#39;basketball&#39;：篮球, 78：&#39;printer&#39;：打印机, 79：&#39;fire hydrant&#39;：消防栓, 80：&#39;projector&#39;：投影仪, 81：&#39;fire extinguisher&#39;：灭火器, 82：&#39;tennis ball&#39;：网球, 83：&#39;frisbee&#39;：飞盘, 84：&#39;fire truck&#39;：消防车, 85：&#39;helicopter&#39;：直升飞机, 86：&#39;carriage&#39;：四轮马车, 87：&#39;bear&#39;：熊, 88：&#39;globe&#39;：地球仪, 89：&#39;volleyball&#39;：排球)。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiObjectdetDetectobjectAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiObjectdetDetectobjectRequest 初始化AliyunViapiObjectdetDetectobjectAPIRequest对象
func NewAliyunViapiObjectdetDetectobjectRequest() *AliyunViapiObjectdetDetectobjectAPIRequest {
	return &AliyunViapiObjectdetDetectobjectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiObjectdetDetectobjectAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.objectdet.detectobject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiObjectdetDetectobjectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiObjectdetDetectobjectAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiObjectdetDetectobjectAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
