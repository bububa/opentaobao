package viapi

// Elements 
type Elements struct {

    // x坐标
    
    X   int64 `json:"x,omitempty" xml:"x,omitempty"`
    

    // y坐标
    
    Y   int64 `json:"y,omitempty" xml:"y,omitempty"`
    

    // 抠图结果（png透明图）有效期半个小时
    
    ImageURL   string `json:"image_u_r_l,omitempty" xml:"image_u_r_l,omitempty"`
    

    // 高度
    
    Height   int64 `json:"height,omitempty" xml:"height,omitempty"`
    

    // 宽度
    
    Width   int64 `json:"width,omitempty" xml:"width,omitempty"`
    

    // 检测框的分数（置信度），范围为[0,1]
    
    Score   int64 `json:"score,omitempty" xml:"score,omitempty"`
    

    // 类型：(0：'human'：人体, 1：'sneakers'：胶底运动鞋, 2：'chair'：椅子, 3：'hat'：帽子, 4：'lamp'：灯, 5：'cabinet/shelf'：橱柜/ 架子, 6：'car'：汽车, 7：'glasses'：眼镜, 8：'picture/frame'：照片/图画, 9：'street lights'：街灯, 10：'helmet'：头盔, 11：'pillow'：枕头, 12：'glove'：手套, 13：'potted plant'：盆栽植物, 14：'flower'：花, 15：'monitor'：显示屏, 16：'plants pot/vase'：花盆, 17：'boots'：靴子, 18：'umbrella'：伞, 19：'boat'：小船, 20：'flag'：旗帜, 21：'speaker'：扬声器/话筒, 22：'trash bin/can'：垃圾桶, 23：'backpack'： 双肩背包, 24：'sofa'：沙发, 25：'belt'：腰带, 26：'carpet'：地毯, 27：'coffee table'：咖啡桌/茶几, 28：'tie'： 领带, 29：'bed'： 床, 30：'traffic light'：红绿灯, 31：'necklace'：项链, 32：'mirror'：镜子, 33：'bicycle'：自行车, 34：'watch'：手表, 35：'horse'：马, 36：'traffic sign'：交通标志, 37：'stuffed animal'：填充玩具动物, 38：'motorbike/motorcycle'：摩托车, 39：'wild bird'：鸟, 40：'laptop'：笔记本电脑, 41：'cow'：奶牛, 42：'clock'：时钟, 43：'bus'：公共汽车, 44：'nightstand'：床头柜, 45：'sheep'：绵羊, 46：'traffic cone'：锥形交通路标, 47：'keyboard'：键盘, 48：'hockey stick'：曲棍球球棍, 49：'fan'：电扇, 50：'dog'：狗, 51：'blackboard/whiteboard'：白板/黑板, 52：'mouse'：鼠标, 53：'telephone'：电话, 54：'airplane'：飞机, 55：'skis'：滑雪板, 56：'soccer'：英式足球, 57：'combine with glove'：棒球手套, 58：'train'：火车, 59：'tent'：帐篷, 60：'sailboat'：帆船, 61：'kite'：风筝, 62：'computer box'：计算机主机机箱, 63：'elephant'：大象, 64：'stroller'：折叠式婴儿车, 65：'baseball bat'：棒球棒, 66：'skateboard'：溜冰板, 67：'surfboard'：冲浪板, 68：'cat'：猫, 69：'zebra'：斑马, 70：'sports car'：跑车, 71：'giraffe'：长颈鹿, 72：'radiator'：散热器, 73：'tennis racket'：网球拍, 74：'skating and skiing shoes'：溜冰鞋, 75：'baseball'：棒球, 76：'american football'：美式橄榄球, 77：'basketball'：篮球, 78：'printer'：打印机, 79：'fire hydrant'：消防栓, 80：'projector'：投影仪, 81：'fire extinguisher'：灭火器, 82：'tennis ball'：网球, 83：'frisbee'：飞盘, 84：'fire truck'：消防车, 85：'helicopter'：直升飞机, 86：'carriage'：四轮马车, 87：'bear'：熊, 88：'globe'：地球仪, 89：'volleyball'：排球)
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 检测框坐标，格式为[left, top, right, bottom]
    
    Boxes   []int64 `json:"boxes,omitempty" xml:"boxes>int64,omitempty"`
    

}
