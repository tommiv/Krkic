// TODO: either build thread-safe singleton
// or do it in the Go way (I wonder if there is the one)

package handlers

var impl []IHandler

func build() {
    impl = append(impl, Basic{})
}

func Impl() []IHandler {
    if len(impl) == 0 {
        build()
    }

    return impl
}
