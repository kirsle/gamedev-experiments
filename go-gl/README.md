# OpenGL Experiments (go-gl)

Playing around with OpenGL for Go.

# Installation

## Go-GL

TBD: Fedora software dependencies, etc. - I probably had most of them installed.

```bash
go get github.com/go-gl/gl/v3.3-core/gl
```

## GLFW3

This one is recommended by the creator of Azul3D Engine:
[Reddit comment](https://www.reddit.com/r/golang/comments/2rnue7/any_opengl_bindings_with_go_that_work_on_windows/cnht2t1)

```bash
# Fedora dependencies... may be others needed too
sudo dnf install libXinerama-devel

go get -u github.com/go-gl/glfw/v3.1/glfw
```

# See Also

* Building on Windows:
  * [Building go-gl/gl on Windows](http://bunkernetz.com/2013/09/01/building-go-glgl-on-windows/)
  * [Windows Build](https://github.com/golang/go/wiki/WindowsBuild)
