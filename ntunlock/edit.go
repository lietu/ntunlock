package ntunlock

type EditSettings struct {
	UnlockAllCharacters bool
	UnlockAllCrowns bool
	UnlockBSkins bool
	UnlockHardMode bool
}

func GetEditSettings() *EditSettings {
	es := EditSettings{
		YesNo("Unlock all characters?") == "y",
		YesNo("Unlock all crowns?") == "y",
		YesNo("Unlock all B-skins?") == "y",
		YesNo("Unlock hard mode?") == "y",
	}

	return &es
}

var AllCrowns []float32 = []float32{
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
	1.000000,
}

func FixCharacter(c *Character, weapon float32) {
	if c.Cwep == 0.0 {
		c.Cwep = weapon
	}

	if c.Startcrown == 0.0 {
		c.Startcrown = 1.0
	}
}

func Edit(save *Save, settings *EditSettings) {
	if settings.UnlockAllCharacters {
		save.Stats.Chicken.Cgot = 1.0
		save.Stats.Crystal.Cgot = 1.0
		save.Stats.Eyes.Cgot = 1.0
		save.Stats.Fish.Cgot = 1.0
		save.Stats.Horror.Cgot = 1.0
		save.Stats.Melting.Cgot = 1.0
		save.Stats.Plant.Cgot = 1.0
		save.Stats.Rebel.Cgot = 1.0
		save.Stats.Robot.Cgot = 1.0
		save.Stats.Rogue.Cgot = 1.0
		save.Stats.Steroids.Cgot = 1.0
		save.Stats.YV.Cgot = 1.0

		// Also set starting weapon and crown for all characters
		FixCharacter(save.Stats.Chicken, 46) // Chicken sword
		FixCharacter(save.Stats.Crystal, 1)
		FixCharacter(save.Stats.Eyes, 1)
		FixCharacter(save.Stats.Fish, 1)
		FixCharacter(save.Stats.Horror, 1)
		FixCharacter(save.Stats.Melting, 1)
		FixCharacter(save.Stats.Plant, 1)
		FixCharacter(save.Stats.Rebel, 1)
		FixCharacter(save.Stats.Robot, 1)
		FixCharacter(save.Stats.Rogue, 81) // Rogue rifle
		FixCharacter(save.Stats.Steroids, 1)
		FixCharacter(save.Stats.YV, 39) // Golden revolver
	}

	if settings.UnlockBSkins {
		save.Stats.Chicken.Cbgt = 1.0
		save.Stats.Crystal.Cbgt = 1.0
		save.Stats.Eyes.Cbgt = 1.0
		save.Stats.Fish.Cbgt = 1.0
		save.Stats.Horror.Cbgt = 1.0
		save.Stats.Melting.Cbgt = 1.0
		save.Stats.Plant.Cbgt = 1.0
		save.Stats.Rebel.Cbgt = 1.0
		save.Stats.Robot.Cbgt = 1.0
		save.Stats.Rogue.Cbgt = 1.0
		save.Stats.Steroids.Cbgt = 1.0
		save.Stats.YV.Cbgt = 1.0
	}

	if settings.UnlockAllCrowns {
		save.Stats.Chicken.Crowns = AllCrowns
		save.Stats.Crystal.Crowns = AllCrowns
		save.Stats.Eyes.Crowns = AllCrowns
		save.Stats.Fish.Crowns = AllCrowns
		save.Stats.Horror.Crowns = AllCrowns
		save.Stats.Melting.Crowns = AllCrowns
		save.Stats.Plant.Crowns = AllCrowns
		save.Stats.Rebel.Crowns = AllCrowns
		save.Stats.Robot.Crowns = AllCrowns
		save.Stats.Rogue.Crowns = AllCrowns
		save.Stats.Steroids.Crowns = AllCrowns
		save.Stats.YV.Crowns = AllCrowns
	}

	if settings.UnlockHardMode {
		save.Data["hardgot"] = 1.0
	}
}