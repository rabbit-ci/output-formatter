require 'formula'

class OutputFormatter < Formula
  homepage 'http://github.com/rabbit-ci/output-formatter'
  version '0.0.1'

  if Hardware.is_64_bit?
    url "https://github.com/rabbit-ci/output-formatter/releases/download/0.0.1/output-formatter_darwin_amd64.zip"
    sha1 '9c604ef7249655f4c47f4e63a5a149bddd4f679c'
  else
    url "https://github.com/rabbit-ci/output-formatter/releases/download/0.0.1/output-formatter_darwin_386.zip"
    sha1 '05a0c9e09e8af16f411e7cc1b9a32a15458f466f'
  end

  def install
    bin.install 'output-formatter'
  end

  test do
    system "output-formatter"
  end
end
