package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

// AuthAvatar implement picture by auth cookie information
type AuthAvatar struct{}

// UseAuthAvatar is a public variable of AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatarURL is the implementation of interface Avatar to type AuthAvatar
func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar implement picture from Gravatar service
type GravatarAvatar struct{}

// UseGravatar is a public variable of GravatarAvatar
var UseGravatar GravatarAvatar

// GetAvatarURL is the implementation of interface Avatar to type GravatarAvatar
func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return fmt.Sprintf("//www.gravatar.com/avatar/%s", useridStr), nil
		}
	}
	return "", ErrNoAvatarURL
}

// FileSystemAvatar implement picture from pc user
type FileSystemAvatar struct{}

// UseFileSystemAvatar is a public variable of FileSystemAvatar
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL is the implementation of interface Avatar to type FileSystemAvatar
func (_ FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					if match, _ := path.Match(useridStr+"*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
