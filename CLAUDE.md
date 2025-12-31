# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

areagen はローグライクゲームのダンジョン生成アルゴリズムを実装したGoプロジェクト。BSP（Binary Space Partitioning）を使用してエリアを再帰的に分割し、部屋と通路を生成してターミナルに描画する。

## コマンド

```bash
# 実行
make run
# または
go run cmd/main.go

# ビルド
go build cmd/main.go

# テスト（テストファイルが追加された場合）
go test ./...
```

## アーキテクチャ

### コア構造体: `Area` (cmd/area/area.go)

`Area`構造体がダンジョン生成の中心。以下のフィールドを持つ：
- `TL`, `BR`: 左上・右下座標（Top-Left, Bottom-Right）
- `Room`: エリア内に生成された部屋（*Area）
- `Child`: BSP分割後の子エリア
- `Path0-3`: 上右下左への通路座標
- `NextTo`: 隣接方向（0=上, 1=右, 2=下, 3=左）

### 生成フロー

1. **Sep()**: 水平/垂直を交互に分割（SepH/SepV）
2. **GenRoom()**: エリア内にランダムサイズの部屋を生成
3. **GenPath()**: 隣接方向へ通路を生成（GenTopPath/GenRightPath/GenBottomPath/GenLeftPath）
4. **LinkPath()**: 親子エリア間の通路を接続
5. **ShowRange()**: ターミナルにANSIカラーで描画

### 座標系

- X軸: 左→右
- Y軸: 上→下
- 座標は`[]int{x, y}`形式で格納
