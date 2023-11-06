package service

import (
	"TinyTik/utils/logger"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 压缩视频
func compressVideo(inputVideoPath string) (string, error) {
	outputVideoPath := strings.TrimSuffix(inputVideoPath, ".mp4") + "CMP.mp4"
	command := []string{
		"-i", inputVideoPath,
		"-c:v", "libx264",
		//"-b:v", "1M", // 使用比特率代替 -crf
		"-crf", "43", //较高的CRF值会减小文件大小但可能降低视频质量
		"-y", // This option enables overwriting without asking
		outputVideoPath,
	}
	cmd := exec.Command("ffmpeg", command...)

	// 打开已存在的日志文件，如果不存在则创建
	logFile, err := os.OpenFile("logs/video.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return outputVideoPath, err
	}
	defer logFile.Close()

	cmd.Stderr = logFile // 将stderr重定向到指定的日志文件
	// cmd.Stderr = os.Stderr // 将stderr重定向到控制台以查看错误消息

	err = cmd.Run()
	if err != nil {
		logger.Debug("Error compressing video:", err)
		return outputVideoPath, err
	}

	return outputVideoPath, nil
}

// 截取封面
func generateVideoCover(videoPath string) (string, error) {
	// 使用 ffmpeg 获取视频的第一帧作为封面
	coverFilename := strings.TrimSuffix(videoPath, ".mp4") + "_cover.jpg"
	command := []string{
		"-i", videoPath,
		"-ss", "00:00:01",
		"-vframes", "1",
		coverFilename,
	}
	cmd := exec.Command("ffmpeg", command...)

	// 打开已存在的日志文件，如果不存在则创建
	logFile, err := os.OpenFile("logs/video.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer logFile.Close()

	cmd.Stderr = logFile // 将stderr重定向到指定的日志文件
	//cmd.Stderr = os.Stderr // Redirect stderr to console for error messages

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating cover:", err)
		return "", err
	}
	return coverFilename, nil
}
