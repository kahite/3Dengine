package main

import (
	"log"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	vertexShaderSource = `#version 330 core
	layout (location = 0) in vec3 aPos;
	
	void main()
	{
		gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
	}`
	fragmentShaderSource = `#version 330 core
	out vec4 FragColor;
	
	void main()
	{
		FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
	}`
)

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, windowsName, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL(orange bool) uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()

	vertexShader := initShaders(gl.VERTEX_SHADER, vertexShaderSource)
	fragmentShader := initShaders(gl.FRAGMENT_SHADER, fragmentShaderSource)

	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return prog
}

func initShaders(shaderType uint32, shaderSource string) uint32 {
	shader := gl.CreateShader(shaderType)
	source, free := gl.Strs(shaderSource)
	gl.ShaderSource(shader, 1, source, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logs := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logs))

		log.Println("failed to compile %v: %v", shaderSource, logs)
	}

	return shader
}
