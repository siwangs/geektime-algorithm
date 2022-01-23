type Twitter struct {
	twm  map[int][]TwitterInfo
	fm   map[int]map[int]struct{}
	time int
}

type TwitterInfo struct {
	t       int
	tweetId int
}

func Constructor() Twitter {
	return Twitter{
		twm:  map[int][]TwitterInfo{},
		fm:   map[int]map[int]struct{}{},
		time: 0,
	}
}

//建立一个新的tweet
func (this *Twitter) PostTweet(userId int, tweetId int) {
	tf := TwitterInfo{t: this.time, tweetId: tweetId}
	this.twm[userId] = append(this.twm[userId], tf)
	this.time++
	if this.fm[userId] == nil {
		this.fm[userId] = map[int]struct{}{userId: {}}
	} else {
		this.fm[userId][userId] = struct{}{}
	}
}

// 检索用户新闻源中的 10 个最新推文 ID。新闻源中的每个项目都必须由用户关注的用户或用户自己发布。推文必须按从最近到最近顺序排序
func (this *Twitter) GetNewsFeed(userId int) []int {
	tfs := make([]TwitterInfo, 0)
	for k, _ := range this.fm[userId] {
		if v, ok := this.twm[k]; ok {
			tfs = append(tfs, v...)
		}
	}
	sort.Slice(tfs, func(i, j int) bool {
		return tfs[i].t > tfs[j].t
	})
	ans := make([]int, 0)
	for i := 0; i < len(tfs) && i < 10; i++ {
		ans = append(ans, tfs[i].tweetId)
	}
	return ans
}

//做关注操作，如果操作无效，不做操作
func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.fm[followerId] == nil {
		this.fm[followerId] = map[int]struct{}{followeeId: {}}
	} else {
		this.fm[followerId][followeeId] = struct{}{}
	}
}

//做取消关注操作，如果操作无效，不做操作
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.fm[followerId] != nil && followeeId != followerId {
		delete(this.fm[followerId], followeeId)
	}
}
    
    