#include <libhbsdcontrol.h>
int
enable_feature(hbsdctrl_ctx_t*, const char*, const char*);
int
disable_feature(hbsdctrl_ctx_t*, const char*, const char*);
int
sysdef_feature(hbsdctrl_ctx_t*, const char*, const char*);
int
feature_status(hbsdctrl_ctx_t*, const char*, const char*, char**);
