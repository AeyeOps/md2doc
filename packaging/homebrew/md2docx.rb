class Md2docx < Formula
  desc "Markdown to Word (.docx) converter"
  homepage "https://github.com/AeyeOps/md2doc"
  url "https://github.com/AeyeOps/md2doc/releases/download/v#{version}/md2docx_#{version}_darwin_amd64.tar.gz"
  sha256 "" # TODO: update with actual SHA256 after release
  license "MIT"

  def install
    bin.install "md2docx"
  end

  test do
    (testpath/"test.md").write("# Hello")
    system "#{bin}/md2docx", "test.md"
    assert_predicate testpath/"test.docx", :exist?
  end
end