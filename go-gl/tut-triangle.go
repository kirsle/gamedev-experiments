package main

// Run this with: `go run tut-triangle.go`
//
// Tutorial used: http://animal-machine.com/blog/141218_getting_started_with_go_and_opengl.md
// But there were some changes to GLFW that required some modifications to the
// tutorial.

import (
	_ "fmt"
	"runtime"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var (
	// Vertex shader. This shader simply copies the vertex location and shares
	// it with the Fragment Shader.
	triangleVertShader = `#version 330
		in vec3 position;
		out vec3 out_pos;
		void main()
		{
			out_pos = position;
			gl_Position = vec4(position, 1.0);
		}`

	// Fragment shader. This shader determines the color per-pixel based off of
	// the fragment location. We also explicitly bind the location for the
	// input position vector for the vertex shader.
	triangleFragShader = `#version 330
		in vec3 out_pos;
		out vec4 colorOut;
		void main()
		{
			float red = out_pos.x + 0.5;
			float green = (-0.5 + out_pos.x) * -1.0;
			float blue = out_pos.y;
			colorOut = vec4(red, green, blue, 1.0);
		}`

	// Shader bind location for the position attribute.
	// shaderAttrLoc_Position = gl.AttribLocation(1)
)

// Loads shader objects and then attaches them to a program.
// func loadShaderProgram() (gl.Program, error) {
// 	// Create the program
// 	prog := gl.CreateProgram()
//
// 	// Create the vertex shader.
// 	vs := gl.CreateShader(gl.VERTEX_SHADER)
// 	vs.Source(triangleVertShader)
// 	vs.Compile()
// 	vsCompiled := vs.Get(gl.COMPILE_STATUS)
// 	if vsCompiled != gl.TRUE {
// 		error := fmt.Sprintf("Failed to compile the vertex shader!\n%s", vs.GetInfoLog())
// 		fmt.Println(error)
// 		return prog, errors.New(error)
// 	}
// 	fmt.Println("Compiled the vertex shader...")
//
// 	return prog
// }

func init() {
	// This is needed to ensure that main() runs on the main thread.
	// See documentation for functions that are only allowed to be called
	// from the main thread.
	runtime.LockOSThread()
}

// Key events are a way to get input from GLFW. Here we check for the escape
// key being pressed. If it's pressed, request that the window be closed.
func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func main() {
	// The GLFW library has to be initialized before any of the methods
	// can be invoked.
	if err := glfw.Init(); err != nil {
		panic("Can't init glfw!")
	}

	// To be tidy, make sure glfw.Terminate() is called at the end of the
	// program to clean things up by using `defer`
	defer glfw.Terminate()

	// Desired number of samples to use for multisampling
	glfw.WindowHint(glfw.Samples, 4)

	// Request an OpenGL 3.3 core context
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	// Create the window
	window, err := glfw.CreateWindow(800, 600, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	gl.Init()

	// Set the callback function to get all key inputs from the user
	window.SetKeyCallback(keyCallback)

	// GLFW3 can work with more than one window, so make sure we set our new
	// window as the current context to operate on.
	window.MakeContextCurrent()

	// Disable v-sync for max FPS if the driver allows it
	glfw.SwapInterval(0)

	// While the window isn't being requested to close (i.e. by the "X" button
	// being clicked, or Escape being pressed)...
	for !window.ShouldClose() {
		// Do OpenGL stuff...

		// swapping OpenGL buffers and polling events has been decoupled in
		// GLFW3, so make sure to invoke both here.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
