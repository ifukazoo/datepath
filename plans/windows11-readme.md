# datepath: Windows 11 対応 README 作成 & レジストリ整備

## Context

- 本リポジトリ (`github.com/yamanobori-old/datepath`) は Go 製ユーティリティ 3 本 (`mkdatedir` / `mkdatememo` / `updatepath`) のセット。
- Windows のエクスプローラー右クリックから日付名ディレクトリ/ファイルを作成する用途で、`right_click.reg` によりレジストリ登録できる。
- README が存在せず、新規ユーザーが Windows 11 で使うための手順が無い。
- 併せて、Windows 11 上で使う際の懸念（右クリックメニュー簡易化、HKLM 権限、PATH 依存、`mkdatememo` レジストリ未登録など）に対応する。

## Windows 11 で使う際の懸念事項と対応方針

| No. | 懸念 | 対応方針 |
| :-- | :-- | :-- |
| 1 | Win11 の右クリックが簡易メニュー化されており、旧式のレジストリ登録項目は「その他のオプションを表示」(Shift+F10) の中に隠れる | README で明記。任意手順として常時詳細メニュー化のレジストリ設定も案内 |
| 2 | 既存 `.reg` が `HKLM` 参照で管理者権限を要求 | `HKCU` 版に置き換え |
| 3 | 既存 `.reg` が `mkdatedir.exe` / `updatepath.exe` を絶対パス無しで参照 → PATH 必須 | `go install` で `%USERPROFILE%\go\bin` に配置し PATH を通す手順を README に記載 |
| 4 | `mkdatememo` のレジストリ登録エントリが欠落 | HKCU 版 `.reg` に追加 |
| 5 | `DateDirUpdate` は `Directory\shell` のみでファイル対象外 | 現行仕様のまま README で明記 |
| 6 | 未署名 exe の SmartScreen 警告 | README のトラブルシューティングで案内 |
| 7 | リポジトリ直下の macOS 用バイナリ `mkdatedir` (2.6MB) | `.gitignore` に `/mkdatedir` を追加 |

## 変更対象ファイル

1. `README.md` — 新規作成（日本語、Windows 11 向けフルスコープ手順）。
2. `right_click.reg` — HKLM 版を削除し HKCU 版に置き換え（UTF-16 LE + BOM、CRLF）。
   - `HKCU\Software\Classes\Directory\Background\shell\DateDirCreate\command` → `"mkdatedir.exe" "%V"`
   - `HKCU\Software\Classes\LibraryFolder\Background\shell\DateDirCreate\command` → `"mkdatedir.exe" "%V"`
   - `HKCU\Software\Classes\Directory\shell\DateDirUpdate\command` → `"updatepath.exe" "%1"`
   - `HKCU\Software\Classes\Directory\Background\shell\DateMemoCreate\command` → `"mkdatememo.exe" "%V"`（新規）
   - `HKCU\Software\Classes\LibraryFolder\Background\shell\DateMemoCreate\command` → `"mkdatememo.exe" "%V"`（新規）
3. `.gitignore` — 末尾に `/mkdatedir` を追加。

Go コードは変更しない。

## README の構成（日本語）

1. **概要** — 3 コマンドの目的と命名規則 `20060102_1504_<name>`。
2. **前提条件** — Windows 11 / Go 1.23 以降（scoop で `scoop install go` 想定）。
3. **セットアップ手順** — `go version` 確認 → clone → `go install ./cmd/...` → PATH 確認。
4. **コマンドラインでの使い方** — 3 コマンドそれぞれの引数と例。
5. **右クリックメニューへの登録** — `right_click.reg` マージ、explorer 再起動、Win11 簡易メニューへの注意。
6. **動作確認** — 実際の右クリック操作の手順。
7. **アンインストール** — `reg delete` コマンドと exe 削除。
8. **トラブルシューティング** — PATH 未通、SmartScreen、メニュー非表示など。

## 検証

- macOS 上で `go build ./...` と `go vet ./...` を実行して Go コードが壊れていないことを確認。
- `.reg` が UTF-16 LE + BOM で保存されているか `file right_click.reg` で確認。
- 実機（Windows 11）での動作確認はユーザー側で実施。
