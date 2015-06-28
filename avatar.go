package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

// AuthAvatar implement picture by auth cookie information
type AuthAvatar struct{}

// TryAvatar to try all avatars possibilities
type TryAvatars []Avatar

// UseAuthAvatar is a public variable of AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatarURL is the implementation of interface Avatar to type AuthAvatar
func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) > 0 {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar implement picture from Gravatar service
type GravatarAvatar struct{}

// UseGravatar is a public variable of GravatarAvatar
var UseGravatar GravatarAvatar

// GetAvatarURL is the implementation of interface Avatar to type GravatarAvatar
func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return fmt.Sprintf("//www.gravatar.com/avatar/%s", u.UniqueID()), nil
}

// FileSystemAvatar implement picture from pc user
type FileSystemAvatar struct{}

// UseFileSystemAvatar is a public variable of FileSystemAvatar
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL is the implementation of interface Avatar to type FileSystemAvatar
func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {

	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}

	return "", ErrNoAvatarURL
}

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}

	return "", ErrNoAvatarURL
}
