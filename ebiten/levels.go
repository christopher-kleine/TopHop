package main

import "github.com/hajimehoshi/ebiten/v2"

type Levels struct {
	Game   *Game
	data   []string
	startX []int
	startY []int
	goalX  []int
	goalY  []int
}

func (l *Levels) Update() error {
	return nil
}

func (l *Levels) Draw(screen *ebiten.Image) {
	l.Game.Player.Draw(screen)
}

func NewLevels(g *Game) *Levels {
	return &Levels{
		Game: g,
		data: []string{
			"eJwTYmBgkKIzFgJiYSBWpDMWHLV31F4smAkNk6pfkkh7iTWfWHXMDJC8xIWGuXH4jZRwgIlxYjGfC2ovOqZmnAvisGMw2Ysr7HkJYGxhSoqduMJAlAAm1W+D2V5sYU/tMB3Fo5hSDACjKTcf",
			"eJwTYmBgkKIzFgJiYSBWpDMWHLV3xNnLRAd7JYm0l4lI9xCrjpkBkpe4sGBuKtjBicNsLqi92DCheCfGXkE85pNqLzY/4Aobcu3lxYKxmUWKWmLsFcWCcdlLrNpReweHvbjyNTXUjmL6YQAQgDcl",
			"eJztk0ELQEAQhSdOFgeLA07+/580slPTZNawEeXwldrXvv3U8wAwPoxHWmR+mObvfXVvJviK73DQWyF1oEDKSJbOKV9FsjlsW3LKnR3SB9Z/MxkcKN8pb3MBD/qO17to67FOa75hfRzNOxXy3OuMeaeieV7xtmzI4nmHt9VT8+bbqBXk5s567nnzbfQKcnMpndzbiW+JzP28gwWWWDpd",
			"eJztk00LgCAMhsVOmR36Olin/v+fzKCBjM3NEC91eEDE7WkvbTbG7I2ZI0vkbMz0ez/ttQk+Mj7nEQF3HtXUmHeNbAJrwbz3LrmHIfM2GHkvj0IvIGUu9dO8obza2bX0TC/Kq5m95L+leoWXXsuctd5O8MKecHlxmUJdmqdD5LywH+l3493FDqquBM7L7W4tb5oNdZfLTcrypx0Xmmk5YA==",
			"eJztlMEKgCAMhkedMjtodrBOvf9LNtAgxA2TsiIP3+1n3/gH0wAwF0YjI7IWRlXv67zNRZmveP9237u9HSKQPrjb8XYSGTySycVmUV4L7pctzL4GmTyGyVGzbGbPqV6KFtyPFhHCbsKe9hy3d0fMFt4bg5unMnKpPOWl+o91dLbLSjk2ZR49nQ==",
			"eJztlLsOgCAMRQlOvgYRB3Xy/3/SmtCkaYpSRCaHk5DK7U1tizPGbJVxwAwcmdhM3fT7ViHF1xIGYAy0QAf0JDaw+6Xq9cASuLQ4lxjzH/3nKdTYBU8EY6m5tL6rud/LXZFH8n3qz1saZb254AxSavjy/mjfScvOvB+pvcG5knJI+6mB6mO+kk7aTw1UL+WO5S3pG/vGZy32Lmi407ufapwqljl0",
			"eJztlLEOwBAURaWdSg20htbU///JkpCIeBW8YOhwBgmOewNJCLk7Iw2H4enMNci7DvJug7xiUu/isL0wA3dAYz8fI6/dx98/5RCJsf5w7sEZeWHP1g29R51ZewZnVIXeFlJe+3dQB0PypHqlATICKzuUD2KUt7bz0l6xstfkm80bdp4jntvi/MHlBS1pOsc=",
			"eJztlDEOgCAMRQlOKgxWGNTJ+1/SkkBCjBQhpjA4vIX89qXQAEKInRlAVuRkZmnslREc3q1yXoVoj6qoHyq9BrEeU1E/vvA+vYG7p7CbRybb41714JXim/2m+vQ0L6dXE7idn5A5cY/xWZylegavJXCZ8I9TvfQtm+tZ4qVyJdmWXvBvkgIiqFxJFn7YuADewjn6",
		},
		startX: []int{1, 1, 1, 1, 1, 1, 1},
		startY: []int{7, 7, 7, 7, 12, 12, 9, 2},
		goalX:  []int{29, 29, 29, 29, 29, 29, 29, 29},
		goalY:  []int{7, 7, 7, 12, 12, 9, 2, 4},
	}
}
