package missevan

/*
 * refer to https://github.com/secriy/kikibot
 */

const (
	TypeRoom     = "room"
	TypeCreator  = "creator"
	TypeGift     = "gift"
	TypeMessage  = "message"
	TypeNotify   = "notify"
	TypeMember   = "member"
	TypeChannel  = "channel"
	TypeQuestion = "question"
	TypeNoble    = "noble"
	TypePK       = "pk"
	TypeSuperFan = "super_fan"
)

// Event defines the message events.
const (
	EventSend         = "send"         // gift send.
	EventNew          = "new"          // new message received.
	EventStatistic    = "statistics"   // statistics of the live room.
	EventJoin         = "join"         // connect to the live room channel.
	EventJoinQueue    = "join_queue"   // members join the live room.
	EventFollowed     = "followed"     // user followed the room creator.
	EventOpen         = "open"         // the live room opened.
	EventClose        = "close"        // the live room closed.
	EventNewRank      = "new_rank"     // the new rank information of the live room.
	EventLeave        = "leave"        // user leaved the live room.
	EventAddAdmin     = "add_admin"    // add a room admin.
	EventRemoveAdmin  = "remove_admin" // Remove a room admin.
	EventRemoveMute   = "remove_mute"  // unmute a user in the live room.
	EventAsk          = "ask"          // ask a question.
	EventAnswer       = "answer"       // answer a question.
	EventConnect      = "connect"      // connect to the live room.
	EventUpdate       = "update"
	EventStart        = "start"
	EventFinish       = "finish"
	EventLastHourRank = "last_hour_rank" // the last hour rank information of the live room.
	EventRenewal      = "renewal"        // 续费（超粉、贵族）事件
	EventRegistration = "registration"   //
	EventHorn         = "horn"           // horn message
)

type SendResp struct {
	Code int `json:"code"`
}

type (
	FMMessage struct {
		Type       string        `json:"type"`
		Event      string        `json:"event"`
		NotifyType string        `json:"notify_type"`
		RoomID     int64         `json:"room_id"`
		Message    string        `json:"message"`
		MessageID  string        `json:"msg_id"`
		User       *FMUser       `json:"user"`
		Queue      []*FMJoin     `json:"queue"`
		Noble      *FMNoble      `json:"noble"`
		SuperFan   *FMSuperFan   `json:"super_fan"`
		Info       *FMInfo       `json:"info"`
		Gift       *FMGift       `json:"gift"`
		Lucky      *FMGift       `json:"lucky"`
		Target     *FMTarget     `json:"target"`
		Statistics *FMStatistics `json:"statistics"`
		PK         *FMPK         `json:"pk"`
		Question   *FMQuestion   `json:"question"`
	}

	// FmUser represents the information of a user.
	FMUser struct {
		IconUrl  string    `json:"iconurl"`
		Titles   []FMTitle `json:"titles"`
		UserID   int64     `json:"user_id"`
		Username string    `json:"username"`
	}

	// FmJoin represents basic information of the user who is joining.
	FMJoin struct {
		Contribution int64     `json:"contribution"`
		IconUrl      string    `json:"iconurl"`
		Titles       []FMTitle `json:"titles"`
		UserID       int64     `json:"user_id"`
		Username     string    `json:"username"`
	}

	FMTitle struct {
		Level int64  `json:"level"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		Color string `json:"color"`
	}

	// FmGift represents the information of gift.
	FMGift struct {
		GiftID       int64  `json:"gift_id"`
		Name         string `json:"name"`
		Price        int64  `json:"price"`
		Number       int64  `json:"num"`
		EffectURL    string `json:"effect_url"`
		WebEffectURL string `json:"web_effect_url"`
	}

	FMTarget struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	FMPK struct {
		ID        string `json:"pk_id"`
		Status    int64  `json:"status"`
		StartTime int64  `json:"start_time"`
		Result    int64  `json:"result"` // PK 结果，0 -> 失败，1 -> 胜利
	}

	FMQuestion struct {
		Username string `json:"username"`
		Question string `json:"question"`
		Price    int64  `json:"price"`
	}

	FMNoble struct {
		Name         string `json:"name"`
		Level        int64  `json:"level"`
		Status       int64  `json:"status"`
		Price        int64  `json:"price"`
		Contribution int64  `json:"contribution"`
	}

	FMSuperFan struct {
		Num int64 `json:"num"`
	}

	FMStatistics struct {
		Accumulation   int64 `json:"accumulation"`    // 累计人数
		Vip            int64 `json:"vip"`             // 贵宾数量
		Score          int64 `json:"score"`           // 分数（热度）
		Revenue        int64 `json:"revenue"`         // 收益
		Online         int64 `json:"online"`          // 在线
		AttentionCount int64 `json:"attention_count"` // 关注数
	}
)

type (
	FMResp struct {
		Success bool    `json:"success"`
		Code    int64   `json:"code"`
		Info    *FMInfo `json:"info"`
	}

	FMInfo struct {
		Creator    *FMCreator `json:"creator"`     // 主播
		Room       *FMRoom    `json:"room"`        // 直播间
		User       *FMUser    `json:"user"`        // 用户
		Medal      *FMMedal   `json:"medal"`       // 粉丝牌
		OwnerCount int64      `json:"owner_count"` //
		Token      string     `json:"token"`
		ExpireAt   int        `json:"expire_at"`
	}

	FMMedal struct {
		CreatorID   int64  `json:"creator_id"`
		CreatorName string `json:"creator_username"`
		Name        string `json:"name"`
		RoomID      int64  `json:"room_id"`
	}

	FMCreator struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	FMRoom struct {
		RoomID       int64         `json:"room_id"`      // 直播间ID
		Name         string        `json:"name"`         // 直播间名
		Announcement string        `json:"announcement"` // 公告
		Members      *FMMembers    `json:"members"`      // 直播间成员
		Statistics   *FMStatistics `json:"statistics"`   // 统计数据
		Status       *FMStatus     `json:"status"`       // 状态信息
		CatalogID    int64         `json:"catalog_id"`   // 子分区ID
		GuildID      int64         `json:"guild_id"`     // 公会ID
		Medal        FMMedal       `json:"medal"`
		Background   struct {
			Enable   bool    `json:"enable"`
			ImageURL string  `json:"image_url"`
			Opacity  float64 `json:"opacity"`
		} `json:"background"` // 背景图
	}

	FMMembers struct {
		Admin []*FMAdmin `json:"admin"` // 管理员
	}

	FMAdmin struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	FMStatus struct {
		Open     int64      `json:"open"`
		OpenTime int64      `json:"open_time"`
		Channel  *FMChannel `json:"channel"`
	}

	FMChannel struct {
		Event    string `json:"event"`
		Platform string `json:"platform"`
		Time     int64  `json:"time"`
		Type     string `json:"type"`
	}

	FMRankMember struct {
		Revenue       int        `json:"revenue"`
		Rank          int        `json:"rank"`
		RankInvisible bool       `json:"rank_invisible"`
		UserId        int        `json:"user_id"`
		Username      string     `json:"username"`
		IconURL       string     `json:"iconurl"`
		Contribution  int        `json:"contribution"`
		Titles        []*FMTitle `json:"titles"`
	}
)
