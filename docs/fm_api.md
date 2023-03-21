# missevan api

## send mesage
request
```
url: https://fm.missevan.com/api/chatroom/message/send
method: post
headers:
    content-type: application/json;charset=UTF-8
    cookie: FM_SESS=20230304|07dm4m7r33k8469kx2x2ngrl7; FM_SESS.sig=QkwDQIwaO5uWsseEHwZmFlNF15w; token=6402a37343c4d6ac29ff1ea1|afa35b312f0ddb95|1677894515|cd4e90e7a0221278

body:
{
    "room_id": 868858631, 
    "message": "你好！",
    "msg_id": "33a286db-6684-48b8-b0fe-dc62b82e44dd",
}
```
response
```
{
    "code": 0,
    "info": {
        "msg_id": "33a286db-6684-48b8-b0fe-dc62b82e44dd",
        "user": {
            "user_id:" 3414815,
            "username": "妳本来就很美",
            "iconurl": "",
            "titles": [
                {
                    "type": "level",
                    "level": 54
                }
            ]
        }
    },
    "noble": null,
    "bubble": null,
    "ok": 1   
}
```

##  chat stream
request
```
url: wss://im.missevan.com/ws?room_id=868858631
method: get
headers:
    Upgrade: websocket
```
connect message
```
{
    "action":"join"
    "uuid":"3d8e8dfe-60e1-4f1c-a54f-77847f1c6048,
    "type":"room",
    "room_id":868858631
}
```
keep alive message
```
❤️
```

## base cookie
request
```
url: https://fm.missevan.com/api/user/info
method: get
```
response
```
headers:
    set-cookie:
```


## login response
```
{"success":true,"code":0,"info":{"id":3414815,"token":"6405e028594838524aeed5c5|7d91945aa6e5de4f|1678106664|8c0c42478e0906c4","user":{"id":3414815,"username":"妳本来就很美","boardiconurl":"icon01.png","iconurl":"https://static.maoercdn.com/avatars/202203/23/57bf93fc43a31c50dd835eac2ad93f65214239.png","userintro":"Someone as handsome as me only appears in anime.","coverurl":"https://static.maoercdn.com/usercover/background.png","avatar":"202203/23/57bf93fc43a31c50dd835eac2ad93f65214239.png","icontype":1,"albumnum":2,"follownum":20,"fansnum":28,"point":318,"userintro_audio":35040,"confirm":2,"soundnum":1,"coverid":0,"avatar2":"https://static.maoercdn.com/avatars/202203/23/57bf93fc43a31c50dd835eac2ad93f65214239.png","boardiconurl2":"https://static.maoercdn.com/profile/icon01.png","duration":3395,"authenticated":0,"hotSound":0,"cvid":0,"soundurl":"https://sound-ks1-cdn-cn.maoercdn.com/aod/202011/02/55628a73636166e1074794cd8fbf0e9f205144.m4a","soundurl_64":"https://sound-ks1-cdn-cn.maoercdn.com/aod/202011/02/55628a73636166e1074794cd8fbf0e9f205144.m4a","soundurl_128":"https://sound-ks1-cdn-cn.maoercdn.com/aod/202011/02/55628a73636166e1074794cd8fbf0e9f205144-128k.m4a","drama_bought_count":0,"balance":444,"teenager_status":0,"operate_type":-1},"expire_at":1709729064}}
```

## FMMessage
```
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"@修电梯的老王 接住小可爱来到直播间\n\n➸ 注音：[xiū diàn tī de lǎo wáng]","msg_id":"3ded2bce-bec8-4c5a-80a9-b07b799e200d"}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"呼呼呼，匿名大佬来一起聊聊天吧ヾ(≧▽≦*)o​​","msg_id":"3f63958e-58e3-4535-987f-c08a0e056d8b"}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"@林タ 接住小耳朵\n\n➸ 注音：[lín]","msg_id":"31fd765f-9fc5-4c4d-8f06-a11458f69edb"}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"@小娇妻_CC可可 接住小可爱进入直播间\n\n➸ 注音：[xiǎo jiāo qī kě kě]","msg_id":"3102e635-da03-4fe7-93e7-26a7cf8a6771"}
{"type":"member","event":"join_queue","notify_type":"","room_id":868904499,"message":"","msg_id":""}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"@王大善人_ 接住小耳朵\n\n➸ 注音：[wáng dà shàn rén]","msg_id":"359cb362-7545-4d32-8db9-b7929fbb564b"}
{"type":"member","event":"join_queue","notify_type":"","room_id":868904499,"message":"","msg_id":""}
{"type":"member","event":"join_queue","notify_type":"","room_id":868904499,"message":"","msg_id":""}
{"type":"notify","event":"horn","notify_type":"message","room_id":115351164,"message":"\u003cfont color=\"#842D11\"\u003e乔爱酥流浪\u003c/font\u003e\u003cfont color=\"#FFFFFF\"\u003e：传下去幺叔淦不动了  他肚子里鹿崽求收养（\u003c/font\u003e\u003cfont color=\"#842D11\"\u003e鹿幺\u003c/font\u003e\u003cfont color=\"#FFFFFF\"\u003e的直播间）\u003c/font\u003e","msg_id":""}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"🍃欢迎小耳朵们点播鸭，当前还没有在排队的项目，发送【点播 XX】来点播吧，( •̀ ω •́ )✧​","msg_id":"3a8a18f5-3bbd-4038-bf38-16b6182de679"}
{"type":"notify","event":"horn","notify_type":"message","room_id":868830033,"message":"\u003cfont color=\"#790202\"\u003e小白酒懵子\u003c/font\u003e\u003cfont color=\"#FFFFFF\"\u003e：连一次白一次，钊钊真的好使🙊🙊（\u003c/font\u003e\u003cfont color=\"#790202\"\u003eMr钊_星禾城\u003c/font\u003e\u003cfont color=\"#FFFFFF\"\u003e的直播间）\u003c/font\u003e","msg_id":""}
{"type":"member","event":"join_queue","notify_type":"","room_id":868904499,"message":"","msg_id":""}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"@Amadeus_DMSk 接住小耳朵","msg_id":"3ab25d95-569d-4506-a773-14af78c52ec0"}
{"type":"message","event":"new","notify_type":"","room_id":868904499,"message":"@休婳婳 \n･*･  ∧,,∧    ∧_∧  ･*･\n'.　 ( ｡･ω･)(･ω･｡ )　.'\n　'･ |   つ♥と       |.･'\n         *ﾟ*ﾟ'  ･｡｡･ﾟ '*\n   * 谢谢小可爱的关注 *\n   * ♬ τнänκ чöü ♬ *\n         ♡阿里嘎多♡​","msg_id":"31a51a65-afb6-4ced-abbb-4c2f8b2aa0f7"}
```