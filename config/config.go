package config

import "sync"

type Config struct {
	filename       string            // 配置文件名字
	lastModifyTime int64             // 上次变更时间
	data           map[string]string //
	rwLock         sync.RWMutex      // 读写锁
	notifyList     []Notifyer        // 切片List
}

type Notifyer interface {
	CallBack(*Config)
}
