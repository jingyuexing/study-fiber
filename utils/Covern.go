/*
* @Author: Jingyuexing
* @Date:   2022-07-02 23:14:05
* @Last Modified by:   Jingyuexing
* @Last Modified time: 2022-07-02 23:18:24
 */
package server_utils

import (
	"unsafe"
)

func Byte2String(s []byte)string{
  return *(*string)(unsafe.Pointer(&s))
}
