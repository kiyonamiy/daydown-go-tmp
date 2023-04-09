# Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/kiyonamiy/daydown.

FROM BASE_IMAGE
LABEL maintainer="<yuqingbo0122@gmail.com>"

WORKDIR /opt/daydown

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/daydown

COPY daydown /opt/daydown/bin/

ENTRYPOINT ["/opt/daydown/bin/daydown"]
CMD ["-c", "/etc/daydown/daydown.yaml"]
