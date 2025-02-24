use crossterm::{
    cursor,
    event::{self, KeyCode},
    execute,
    terminal::{self, ClearType},
    ExecutableCommand,
};

use std::io::{stdout, Write};

pub struct Nevara {
    pub width: u16,
    pub height: u16
}