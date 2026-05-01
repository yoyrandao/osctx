Output docs/img/demo.gif

Require osctx
Require fzf

Set Shell "bash"
Set FontSize 15
Set Width 750
Set Height 380
Set Theme "Catppuccin Mocha"
Set Padding 16
Set TypingSpeed 60ms
Set PlaybackSpeed 1.5

# ── hidden setup ────────────────────────────────────────────────────────────
Hide

Type "mkdir -p /tmp/osctx-demo"
Enter
Sleep 100ms

Type `cat > /tmp/osctx-demo/clouds.yaml << 'EOF'`
Enter
Type "clouds:"
Enter
Type "  dev:"
Enter
Type "    auth:"
Enter
Type "      auth_url: https://dev.openstack.example.com:5000"
Enter
Type "  staging:"
Enter
Type "    auth:"
Enter
Type "      auth_url: https://staging.openstack.example.com:5000"
Enter
Type "  production:"
Enter
Type "    auth:"
Enter
Type "      auth_url: https://prod.openstack.example.com:5000"
Enter
Type "  us-east-1:"
Enter
Type "    auth:"
Enter
Type "      auth_url: https://us-east.openstack.example.com:5000"
Enter
Type "EOF"
Enter
Sleep 200ms

Type "cd /tmp/osctx-demo"
Enter

Type `osctx() { eval "$(command osctx "$@")"; }`
Enter
Sleep 100ms

Show

# ── demo ────────────────────────────────────────────────────────────────────

Type "osctx ls"
Sleep 500ms
Enter
Sleep 1s

Type "osctx current"
Sleep 500ms
Enter
Sleep 1s

Type "osctx"
Sleep 500ms
Enter
Sleep 800ms

# navigate fzf: move down to "production"
Down
Sleep 400ms
Down
Sleep 400ms
Enter
Sleep 800ms

Type "osctx current"
Sleep 500ms
Enter
Sleep 1s

Type "osctx unset"
Sleep 500ms
Enter
Sleep 1s

Type "osctx current"
Sleep 500ms
Enter
Sleep 1.5s
