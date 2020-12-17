use crate::pane::Pane;
use crate::pane::PaneType;

use tui::layout::{Constraint, Direction, Layout};
use std::process::Command;
use std::str;
use std::error::Error;

struct Window {
    count: u8,
}

impl Window {
    fn new(count: u8) -> Window {
        Window {
            count,
        }
    }
}

fn get_rect(count: u8, pane: PaneType, layouts: &Vec<tui::layout::Rect>) -> tui::layout::Rect {
    match pane {
        PaneType::Parent => layouts[0],
        PaneType::Current => {
            if count == 2{
                return layouts[3];
            }
            layouts[1]
        },
        PaneType::Child => layouts[2],
    }
}

pub fn draw(terminal: &mut tui::Terminal<tui::backend::TermionBackend<termion::raw::RawTerminal<std::io::Stdout>>>) -> Result<(), Box<dyn Error>>{
    terminal.draw(|t| {
        let mut win = Window::new(0);

        let chunks = Layout::default()
            .direction(Direction::Horizontal)
            .constraints(
                [
                    Constraint::Percentage(20),
                    Constraint::Percentage(50),
                    Constraint::Percentage(30),
                    Constraint::Percentage(70),
                ]
                .as_ref(),
            )
            .split(t.size());

        //parent pane 
        let parent_ls = Command::new("ls").arg("..").output().unwrap();
        win.count = 3;
        //idk how to handle Result<T, E> at all
        if win.count == 3 {
            let parent_files = str::from_utf8(&parent_ls.stdout).unwrap();
            let parent_pane = Pane::new(parent_files, PaneType::Parent); 
            parent_pane.visualize(t, get_rect(win.count, PaneType::Parent, &chunks));
        }
        
        //current pane
        let curr_ls = Command::new("ls").output()
            .expect("Unable to read directory");
        let curr_files = str::from_utf8(&curr_ls.stdout).unwrap();
        
        let main_pane = Pane::new(curr_files, PaneType::Current);
        main_pane.visualize(t, get_rect(win.count, PaneType::Current, &chunks));

        //child pane
        //temp command
        let child_ls = Command::new("ls").arg("-a").output() 
            .expect("Error reading directory");
        let child_files = str::from_utf8(&child_ls.stdout).unwrap();

        let child_pane = Pane::new(child_files, PaneType::Child);
        child_pane.visualize(t, get_rect(win.count, PaneType::Child, &chunks));

    })?;
    Ok(())
}
