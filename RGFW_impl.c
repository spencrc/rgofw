#define RGFW_IMPLEMENTATION
#include "RGFW_impl.h"

wrapper_eventData wrapper_getEventData(RGFW_event* e) {
    wrapper_eventData d = {
        e->type,
        e->button.value,
        e->scroll.x,
        e->scroll.y,
        e->mouse.x,
        e->mouse.y,
        e->mouse.vecX,
        e->mouse.vecY,
        e->key.value,
        e->key.repeat,
        e->key.mod,
        e->keyChar.value,
        e->drop.files,
        e->drop.count,
        e->drag.x,
        e->drag.y,
        e->scale.x,
        e->scale.y,
        e->monitor.monitor
    };
    return d;
}