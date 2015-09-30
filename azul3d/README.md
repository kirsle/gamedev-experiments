# Azul3D Experiments

Playing around with the [Azul3D](https://azul3d.org) Golang game engine.

There's not much documentation or very good examples for Azul3D yet... so I
only got as far as being able to run the examples.

# Installation

## Fedora Dependencies

The [official docs](https://azul3d.org/doc/install/linux.html) have dependency
lists for Ubuntu; these are what I found is needed for Fedora (as of Fedora 22).
There might be some others; I installed these on my main system which may have
had other dependencies installed already (e.g. `gcc` or whatever).

```bash
sudo dnf install git libX11-devel libxcb-devel libXrandr-devel \
    libXcursor-devel libXi-devel freetype-devel
```

## Azul3D

The docs said to do e.g. `go get -u azul3d.org/gfx.v1` but that URL gave a 404.
Additionally, their GitHub links refer to the other modules by the `azul3d.org`
names, so they couldn't resolve their dependencies either.

Workaround: check out the modules manually into `$GOPATH/src/azul3d.org` for
example:

```bash
mkdir -p ~/go/src/azul3d.org
cd ~/go/src/azul3d.org
git clone --branch v1.0.1 https://github.com/azul3d/gfx gfx.v1
git clone --branch v1 https://github.com/azul3d/audio audio.v1
git clone --branch v2 https://github.com/azul3d/gfx-window gfx/window.v2
git clone --branch v1 https://github.com/azul3d/lmath lmath.v1
git clone --branch v1 https://github.com/azul3d/clock clock.v1
git clone --branch v2 https://github.com/azul3d/gfx-gl2 gfx/gl2.v2
git clone --branch v1 https://github.com/azul3d/mouse mouse.v1
git clone --branch v3.1 https://github.com/azul3d/native-glfw native/glfw.v3.1
```
