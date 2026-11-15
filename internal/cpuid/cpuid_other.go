//go:build !amd64 || appengine || noasm || !gc || purego

package cpuid

func detect() Info {
	return Info{}
}
