# Add Architecture Image to GitHub - Step by Step

## What You Need

The architecture image you showed me earlier. This should be saved as a PNG or JPG file on your computer.

---

## Method 1: Via GitHub Web UI (Easiest, No Command Line)

### Step 1: Save the Image
1. Right-click the architecture image on your computer
2. Save it as: `architecture.png`
3. Remember where you saved it

### Step 2: Upload via GitHub Website
1. Go to: https://github.com/izhan05803/lsm-tree-database
2. Click **"Add file"** button (top right, near "Code")
3. Select **"Upload files"**
4. **Drag and drop** your `architecture.png` file
5. OR click **"choose your files"** and select it

### Step 3: Create the Folder Path
GitHub will ask where to save the file. Type:
```
docs/images/architecture.png
```

### Step 4: Commit
- Add commit message: `"Add architecture diagram image"`
- Click **"Commit changes"**

Done! The image is now on GitHub at `docs/images/architecture.png`

---

## Method 2: Via Command Line (More Control)

### Step 1: Create Folder
```bash
cd D:\embedded-lsm-tree-engine
mkdir -p docs/images
```

### Step 2: Add Your Image
Copy your architecture image file to:
```
D:\embedded-lsm-tree-engine\docs\images\architecture.png
```

### Step 3: Push to GitHub
```bash
git add docs/
git commit -m "Add architecture diagram image"
git push origin main
```

Done! The image is now on GitHub.

---

## Verify It Worked

1. Go to: https://github.com/izhan05803/lsm-tree-database
2. You should see a `docs` folder in the file list
3. Inside it, see `images` folder
4. Inside that, see `architecture.png`
5. Look at the README - the image should display!

If the image doesn't show in README immediately:
- GitHub sometimes takes 5-10 seconds to refresh
- Try hard refresh: **Ctrl + Shift + R** (or Cmd + Shift + R on Mac)

---

## Image Best Practices

### File Size
- Keep under 1MB for fast loading
- Recommended: 600-800px wide

### Format
- PNG: Best for diagrams and screenshots
- JPG: Best for photos

### File Name
- Use lowercase: `architecture.png` ✅
- Not: `Architecture.png` or `arch_diagram_v2.png`
- Use hyphens: `architecture-diagram.png`

---

## Troubleshooting

### Image doesn't show up in README
1. Make sure file path is correct: `docs/images/architecture.png`
2. Check file extension matches: `.png` or `.jpg`
3. Hard refresh GitHub: **Ctrl + Shift + R**

### Git error when pushing
```
error: failed to push some refs
```
Solution:
```bash
git pull origin main
git push origin main
```

### File is too large
```bash
# Compress the image using ImageMagick (if installed)
convert architecture.png -resize 800x600 architecture-small.png
```

---

## Final Directory Structure

After adding the image, your repo should look like:

```
lsm-tree-database/
├── README.md                           ← References the image
├── docs/
│   └── images/
│       └── architecture.png            ← Your architecture diagram
├── memtable.go
├── wal.go
├── sstable.go
├── db.go
├── main.go
├── go.mod
└── .gitignore
```

---

## Next Steps

1. **Save your image** as `architecture.png`
2. **Choose Method 1 or 2** above
3. **Verify** it shows on GitHub
4. Done! Your README now has a visual architecture diagram

---

## If You Need Multiple Images Later

For write flow, read flow, etc.:

```
docs/images/
├── architecture.png
├── write-flow.png
├── read-flow.png
└── recovery-flow.png
```

Reference in README:
```markdown
![Write Flow](docs/images/write-flow.png)
![Read Flow](docs/images/read-flow.png)
```

---

## Questions?

Check the main README at:
https://github.com/izhan05803/lsm-tree-database
