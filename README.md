# datepath

Windows のエクスプローラーから、日時プレフィックス付きのディレクトリ／メモファイルを
右クリック 1 回で作成・更新するための小さな Go 製ユーティリティ集です。

## 提供コマンド

| コマンド | 役割 | 生成／変更されるパス |
| :-- | :-- | :-- |
| `mkdatedir <dir> [name]` | `<dir>` の下に日時付きディレクトリを作成 | `20060102_1504_<name>/` |
| `mkdatememo <dir> [name]` | `<dir>` の下に日時付きの空メモファイルを作成 | `20060102_1504_<name>.txt` |
| `updatepath <path>` | 既に日時プレフィックスが付いているパスの日時部分を現在時刻に更新 | 例: `20240101_1200_memo.txt` → `20260720_1830_memo.txt` |

`[name]` を省略した場合は日時のみのパス（末尾のアンダースコアが残る形）になります。

日時フォーマットは `YYYYMMDD_HHMM` 固定です（`common/common.go` の `DateFormat` を参照）。

## 動作環境

- Windows 11（Windows 10 でも同じ手順で動作します）
- Go 1.23 以降

## セットアップ

### 1. Go をインストールする

[scoop](https://scoop.sh/) 経由でのインストールを想定しています。

```powershell
scoop install go
go version
```

`go version go1.23.x windows/amd64` のように表示されれば OK です。

### 2. リポジトリを取得してビルド・インストールする

```powershell
git clone https://github.com/ifukazoo/datepath.git
cd datepath
go install ./cmd/mkdatedir ./cmd/mkdatememo ./cmd/updatepath
```

`go install` により 3 つの exe が `%USERPROFILE%\go\bin\` に配置されます。
右クリックメニューは同梱 `right_click.reg` がこのパスを `%USERPROFILE%` で
参照するため、PATH の追加は不要です。

### 3. （任意）PATH を通す — コマンドラインから直接叩きたい場合

コマンドプロンプトや PowerShell で `mkdatedir` などを直接タイプして使いたい場合は、
`%USERPROFILE%\go\bin` を User PATH に追加します。右クリックメニューだけ使う場合は
この節はスキップして構いません。

- GUI 手順: `Windows キー` → 「環境変数を編集」→ ユーザー環境変数の `Path` → 「編集」→ 「新規」→ `%USERPROFILE%\go\bin` を追加 → 開いている PowerShell を再起動。
- コマンド 1 行で追加する場合:
  ```powershell
  [Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$Env:USERPROFILE\go\bin", "User")
  ```
  こちらも設定後は PowerShell の再起動が必要です。

反映確認:

```powershell
where mkdatedir
```

`C:\Users\<あなた>\go\bin\mkdatedir.exe` が返れば通っています。

## コマンドラインでの使い方

いずれのコマンドも第 1 引数の存在しないパスを渡すとエラーになります。

```powershell
mkdatedir  C:\work           作業           # → C:\work\20260720_1830_作業\
mkdatememo C:\work           議事メモ       # → C:\work\20260720_1830_議事メモ.txt
updatepath C:\work\20240101_1200_作業       # → C:\work\20260720_1830_作業\ にリネーム
```

`updatepath` は先頭 13 文字が `YYYYMMDD_HHMM` の形式になっているパスにのみ使えます。

## 右クリックメニューへの登録

同梱の `right_click.reg` を使うと、エクスプローラーの右クリックから直接呼び出せるようになります。
このファイルは `HKEY_CURRENT_USER` に書き込むため、管理者権限は不要です。
コマンドの実行パスは `%USERPROFILE%\go\bin\<exe>` を `REG_EXPAND_SZ` として登録するため、
Windows がログオンユーザーに応じて自動で展開します。ユーザー名の書き換えや PATH 追加は不要です。

### 登録

1. `right_click.reg` をダブルクリック → 「はい」で結合を承認。
2. 反映のため、開いているエクスプローラーを一度閉じるか、次のコマンドでエクスプローラーを再起動:
   ```powershell
   taskkill /f /im explorer.exe; start explorer
   ```

### 登録される項目

| 右クリック対象 | メニュー名 | 実行内容 |
| :-- | :-- | :-- |
| フォルダの何もない所（背景） | `date dir create` | 現在のフォルダに日時付きディレクトリを作成 |
| フォルダの何もない所（背景） | `date memo create` | 現在のフォルダに日時付きメモを作成 |
| 既存のフォルダ | `date dir update` | フォルダ名の日時プレフィックスを現在時刻に更新 |

「ライブラリ」内（ドキュメント／ピクチャ等）にも同じ項目が追加されます。

### Windows 11 での注意

Windows 11 の既定のコンテキストメニューは簡易表示のため、追加された項目は
**「その他のオプションを表示」（Shift+F10 でも可）** の下に隠れます。

常に旧来の詳細メニューを表示させたい場合は、PowerShell で以下を実行してください
（要ログアウト or エクスプローラー再起動）。取り消したいときはこの下のアンインストール
に手順を書いています。

```powershell
reg add "HKCU\Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}\InprocServer32" /f /ve
```

## 動作確認

1. 適当な作業フォルダを開く。
2. フォルダの何もない場所を右クリック → `date dir create` → `YYYYMMDD_HHMM_` という名前のフォルダが生成される。
3. 続けて `date memo create` → 同様の名前で `.txt` が生成される。
4. 生成された日時付きフォルダを右クリック → `date dir update` → 日時部分だけが現在時刻に置き換わる。

コマンド名は決め打ちのため、生成後に `<name>` 部分を付けたい場合はエクスプローラー上でリネームしてください。

## アンインストール

### レジストリを削除

PowerShell（管理者不要）で以下を実行してください。

```powershell
reg delete "HKCU\Software\Classes\Directory\Background\shell\DateDirCreate"    /f
reg delete "HKCU\Software\Classes\LibraryFolder\Background\shell\DateDirCreate" /f
reg delete "HKCU\Software\Classes\Directory\Background\shell\DateMemoCreate"   /f
reg delete "HKCU\Software\Classes\LibraryFolder\Background\shell\DateMemoCreate" /f
reg delete "HKCU\Software\Classes\Directory\shell\DateDirUpdate"                /f
```

Windows 11 の「常に詳細メニュー」設定も戻したい場合:

```powershell
reg delete "HKCU\Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}" /f
```

### exe を削除

```powershell
Remove-Item "$Env:USERPROFILE\go\bin\mkdatedir.exe"
Remove-Item "$Env:USERPROFILE\go\bin\mkdatememo.exe"
Remove-Item "$Env:USERPROFILE\go\bin\updatepath.exe"
```

## トラブルシューティング

- **右クリックメニューに項目が出ない** — Windows 11 では簡易メニューに隠れるため、「その他のオプションを表示」または `Shift+F10` で開いてください。それでも出ない場合は `explorer.exe` を再起動。
- **右クリックしても何も起きない／一瞬コンソールが開いて消える** — `%USERPROFILE%\go\bin\` に exe が入っているか確認してください。存在しない場合は `go install ./cmd/mkdatedir ./cmd/mkdatememo ./cmd/updatepath` を再実行。
- **`where mkdatedir` で見つからない**（コマンドラインから使いたい場合のみ） — `%USERPROFILE%\go\bin` を PATH に追加し、PowerShell を再起動してください（上記「セットアップ 3」）。右クリック用途では PATH は関係しません。
- **SmartScreen で `mkdatedir.exe` の起動がブロックされる** — `go install` でローカルビルドした exe は通常ブロックされませんが、警告が出たら「詳細情報」→「実行」で許可してください。
- **`updatepath` が `can't update path` エラー** — 対象パスの先頭が `YYYYMMDD_HHMM` の形式でないと更新できません。日時プレフィックスの有無を確認してください。
- **日本語の `<name>` が化ける** — 生じない想定です。もし化けた場合はコンソールのコードページ（`chcp`）を確認してください。

## ライセンス

[LICENSE](LICENSE) を参照。
