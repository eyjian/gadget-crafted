#!/bin/sh

HOME=/app/gadget-crafted
cd $HOME
if test $? -eq 0; then
    $HOME/gadget_crafted_service > $HOME/gadget_crafted_service.nohup 2>&1 &
fi
