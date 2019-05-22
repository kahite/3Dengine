package main

import (
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width       = 1000
	height      = 1000
	windowsName = "Hello Window"
)

var vertices = []float32{
	0.5, 0.5, 0.0,
	0.5, -0.5, 0.0,
	-0.5, -0.5, 0.0,
	-0.5, 0.5, 0.0,
}

var indices = []int32{
	0, 1, 3,
	1, 2, 3,
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL(true)

	var dummy float32
	var vbo, vao, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(dummy)), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(unsafe.Sizeof(ebo)), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*int32(unsafe.Sizeof(dummy)), nil)
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		// gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
		gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)
		gl.BindVertexArray(0)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}

	if window.GetKey(glfw.KeyQ) == glfw.Press {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}

	if window.GetKey(glfw.KeyW) == glfw.Press {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}
}
